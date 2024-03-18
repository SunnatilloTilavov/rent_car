package postgres

import (
	"database/sql"
	"fmt"
	"clone/rent_car_us/models"
	"clone/rent_car_us/pkg"
	"github.com/google/uuid"
)

type customerRepo struct {
	db *sql.DB
}

func NewCustomer(db *sql.DB) customerRepo {
	return customerRepo{
		db: db,
	}
}

/*
create (body) id,err
update (body) id,err
delete (id) err
get (id) body,err
getAll (search) []body,count,err
*/

func (c *customerRepo) Create(customer models.Customer) (string, error) {

	id := uuid.New()

	query := ` INSERT INTO customers (
		id ,       
		first_name,
		last_name ,
		gmail ,    
		phone     )
		VALUES($1,$2,$3,$4,$5) `

	_, err := c.db.Exec(query,
		id.String(),
		customer.First_name, customer.Last_name,
		customer.Gmail, customer.Phone)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (c *customerRepo) Update(customer models.Customer) (string, error) {

	query := ` UPDATE customers set
	        first_name=$1,
	        last_name=$2,
	        gmail=$3,
	        phone=$4,
			updated_at=CURRENT_TIMESTAMP
		WHERE id = $5 AND deleted_at=0
	`

	_, err := c.db.Exec(query,
		customer.First_name, customer.Last_name,
		customer.Gmail, customer.Phone,customer.Id)

	if err != nil {
		return "", err
	}

	return customer.Id, nil
}

func (c *customerRepo) GetAllCustomers(req models.GetAllCustomersRequest) (models.GetAllCustomersResponse, error) {
	var (
		resp   = models.GetAllCustomersResponse{}
		filter = ""
	)
	offset := (req.Page - 1) * req.Limit
	if req.Search != "" {
		filter += fmt.Sprintf(`and first_name ILIKE '%%%v%%'`, req.Search)
	}

	filter += fmt.Sprintf("OFFSET %v LIMIT %v", offset, req.Limit)
	fmt.Println("filter:", filter)

	query:=`Select 
	cu.id as customer_id,
	cu.first_name as customer_first_name,
	cu.last_name as customer_last_name,
	cu.gmail as customer_gmail,
	cu.phone as customer_phone,
	cu.created_at,
	cu.updated_at,
	o.id,
	o.from_date,
	o.to_date,
	o.status,
	o.paid,
	o.amount
	From customers cu JOIN orders o ON  cu.id = o.customer_id`

	rows, err := c.db.Query(query + filter + ``)
	if err != nil {
		return resp, err
	}
     defer rows.Close()

	for rows.Next() {
		var (
			customer = models.GetAllCustomer{
				Order: models.GetOrder{},
			}
			updateAt sql.NullString
		)
		if err := rows.Scan(
			&customer.Id,
			&customer.FirstName,
			&customer.LastName,
			&customer.Gmail,
			&customer.Phone,
			&customer.Is_Blocked,
			&customer.CreatedAt,
			&updateAt,
			&customer.Order.Id,       
			&customer.Order.FromDate,       
			&customer.Order.ToDate,       
			&customer.Order.Status,       
			&customer.Order.Paid,       
			&customer.Order.Amount);err != nil {
			return resp, err
		}
		customer.UpdatedAt = pkg.NullStringToString(updateAt)
		resp.Customers = append(resp.Customers, models.Customer{})
	}
	if err = rows.Err();err != nil {
		return resp,err
	}
	countQuery := `Select count(*) from customers`
	err = c.db.QueryRow(countQuery).Scan(&resp.Count)
	if err != nil {
		return resp,err
	}
	return resp, nil
}

func (c *customerRepo) GetByID(id string) (models.Customer, error) {
	customer := models.Customer{}

	if err := c.db.QueryRow(`select id,first_name,last_name,gmail,phone,is_blocked from customers where id = $1`, id).Scan(
		&customer.Id,
		&customer.First_name,
		&customer.Last_name,
		&customer.Gmail,
		&customer.Phone); err != nil {
		return models.Customer{}, err
	}
	return customer, nil
}
func (c *customerRepo) Delete(id string) error {

	query := ` UPDATE customers set
			deleted_at = date_part('epoch', CURRENT_TIMESTAMP)::int
		WHERE id = $1 AND deleted_at=0
	`

	_, err := c.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

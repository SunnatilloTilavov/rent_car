package postgres


import (
	"clone/rent_car_us/config"
	"clone/rent_car_us/storage"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Store struct {
	DB       *sql.DB
}
func New(cfg config.Config) (storage.IStorage, error) {
	url := fmt.Sprintf(`host=%s port=%v user=%s password=%s database=%s sslmode=disable`,
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabase)

	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return Store{
		DB: db,
	}, nil
}
func (s Store) CloseDB() {
	s.DB.Close()
}

func (s Store) Car() storage.ICarStorage {
	newCar := NewCar(s.DB)

	return &newCar
}
func (s Store) Customer() storage.ICustomerStorage {
	newCustomer := NewCustomer(s.DB)

	return &newCustomer
}
func (s Store) Order() storage.IOrderStorage {
	newOrder := NewOrder(s.DB)

	return &newOrder
}




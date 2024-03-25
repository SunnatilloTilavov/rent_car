package handler

import (
	"fmt"
	_ "clone/rent_car_us/api/docs"
	"clone/rent_car_us/api/models"
	"clone/rent_car_us/pkg/check"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)
// CreateCustomer godoc
// @Router 		/customer [POST]
// @Summary 	create a customer
// @Description This api is creates a new customer and returns it's id
// @Tags 		customer
// @Accept		json
// @Produce		json
// @Param		customer body models.Customer true "customer"
// @Success		200  {object}  models.Customer
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateCustomer(c *gin.Context) {
	Customer := models.Customer{}

	if err := c.ShouldBindJSON(&Customer); err != nil {
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.Store.Customer().Create(Customer)
	if err != nil {
		handleResponse(c, "error while creating Customer", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, "Created successfully", http.StatusOK, id)
}
// Updatecustomer godoc
// @Router 		/customer/{id} [PUT]
// @Summary 	update a customer
// @Description This api is update a  customer and returns it's id
// @Tags 		customer
// @Accept		json
// @Produce		json
// @Param		id path string true "id"
// @Param		customer body models.Customer true "customer"
// @Success		200  {object}  models.Customer
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) UpdateCustomer(c *gin.Context) {
	Customer := models.Customer{}

	if err := c.ShouldBindJSON(&Customer); err != nil {
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	if err := check.ValidateEmail(Customer.Gmail); err != nil {
		handleResponse(c, "error while validating Customer year, year: "+Customer.Gmail, http.StatusBadRequest,err.Error())
		return
	}

	if err := check.ValidatePhone(Customer.Phone); err != nil {
		handleResponse(c, "error while validating Customer year, year: "+Customer.Phone, http.StatusBadRequest,err.Error())
		return
	}
	Customer.Id = c.Param("id")

	err := uuid.Validate(Customer.Id)
	if err != nil {
		handleResponse(c, "error while validating Customer id,id: "+Customer.Id, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Store.Customer().Update(Customer)
	if err != nil {
		handleResponse(c, "error while updating Customer", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, "Updated successfully", http.StatusOK, id)
}
// GETALLCustomerS godoc
// @Router 		/customer [GET]
// @Summary 	Get customer list
// @Description Get customer list
// @Tags 		customer
// @Accept		json
// @Produce		json
// @Param		page path string false "page"
// @Param		limit path string false "limit"
// @Param		search path string false "search"
// @Success		200  {object}  models.Customer
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response

func (h Handler) GetAllCustomers(c *gin.Context) {
	var (
		request = models.GetAllCustomersRequest{}
	)

	request.Search = c.Query("search")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponse(c, "error while parsing page", http.StatusBadRequest, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleResponse(c, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("page: ", page)
	fmt.Println("limit: ", limit)

	request.Page = page
	request.Limit = limit
	Customers, err := h.Store.Customer().GetAllCustomers(request)
	if err != nil {
		handleResponse(c, "error while gettign Customers", http.StatusBadRequest, err.Error())

		return
	}

	handleResponse(c, "", http.StatusOK, Customers)
}
// Deletecustomer godoc
// @Router 		/customer/{id} [DELETE]
// @Summary 	delete a customer
// @Description This api is delete a  customer and returns it's id
// @Tags 		customer
// @Accept		json
// @Produce		json
// @Param		id path string true "id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response


func (h Handler) DeleteCustomer(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id: ", id)

	err := uuid.Validate(id)
	if err != nil {
		handleResponse(c, "error while validating id", http.StatusBadRequest, err.Error())
		return
	}

	err = h.Store.Customer().Delete(id)
	if err != nil {
		handleResponse(c, "error while deleting Customer", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, id)
}

// GETBYIDCustomer godoc
// @Router 		/customer [GET]
// @Summary 	Get customer 
// @Description Get customer
// @Tags 		customer
// @Accept		json
// @Produce		json
// @Param		id path string true "id"
// @Success		200  {object}  models.Customer
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response

func (h Handler) GetByIDCustomer(c *gin.Context) {
 
	id := c.Param("id")
	fmt.Println("id: ", id)
   
	admin, err := h.Store.Customer().GetByID(id)
	if err != nil {
	 handleResponse(c, "error while getting admin by id", http.StatusInternalServerError, err)
	 return
	}
	handleResponse(c, "", http.StatusOK, admin)
   }



   func (h Handler) GetAllCustomerCars(c *gin.Context) {
	var (
		request = models.GetAllCustomerCarsRequest{}
	)

	request.Search = c.Query("search")
	request.Id=c.Param("id")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponse(c, "error while parsing page", http.StatusBadRequest, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleResponse(c, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("page: ", page)
	fmt.Println("limit: ", limit)

	request.Page = page
	request.Limit = limit
	Orders, err := h.Store.Customer().GetAllCustomerCars(request)
	if err != nil {
		handleResponse(c, "error while gettign CustomerCars", http.StatusBadRequest, err.Error())

		return
	}

	handleResponse(c, "", http.StatusOK, Orders)
}
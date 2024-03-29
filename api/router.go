package api

import (
	"clone/rent_car_us/api/handler"
	"clone/rent_car_us/service"
	"clone/rent_car_us/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New ...
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
func New(services service.IServiceManager,store storage.IStorage) *gin.Engine {
	h := handler.NewStrg(store,services)

	r := gin.Default()

	r.POST("/car", h.CreateCar)
	r.GET("/car/:id", h.GetByIDCar)
	r.GET("/car", h.GetAllCars)
	r.GET("/car/free", h.GetAllCarsFree)
	r.PUT("/car/:id", h.UpdateCar)
	r.DELETE("/car/:id", h.DeleteCar)
	// r.PATCH("/car/:id", h.UpdateUserPassword)

	r.POST("/customer", h.CreateCustomer)
	r.GET("/customer/:id", h.GetByIDCustomer)
	r.GET("/customer", h.GetAllCustomers)
	r.PUT("/customer/:id", h.UpdateCustomer)
	r.DELETE("/customer/:id", h.DeleteCustomer)
	r.GET("/customercars",h.GetAllCustomerCars)

	r.POST("/order", h.CreateOrder)
	r.GET("/order/:id", h.GetOne)
	r.GET("/order", h.GetAllOrders)
	r.PUT("/order/:id", h.UpdateOrder)
	r.DELETE("/order/:id", h.DeleteOrder)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
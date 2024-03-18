package storage

import "clone/rent_car_us/models"

type IStorage interface {
	CloseDB()
	Car() ICarStorage
	Customer() ICustomerStorage
	Order() IOrderStorage
}

type ICarStorage interface {
	Create(models.Car) (string, error)
	GetByID(id string) (models.Car, error)
	GetAllCars(request models.GetAllCarsRequest) (models.GetAllCarsResponse, error)
	Update(models.Car) (string, error)
	Delete(string) error
}

type ICustomerStorage interface {
	Create(models.Customer) (string, error)
	GetByID(id string) (models.Customer, error)
	GetAllCustomers(request models.GetAllCustomersRequest) (models.GetAllCustomersResponse, error)
	Update(models.Customer) (string, error)
	Delete(string) error
}

type IOrderStorage interface {
	CreateOrder(models.CreateOrder) (string, error)
	UpdateOrder(models.GetOrder) (string, error)
	GetOne(id string) (models.GetOrder, error)
	GetAll(request models.GetAllOrdersRequest) (models.GetAllOrdersResponse, error)
	DeleteOrder(string) error
}
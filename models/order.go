package models

type GetOrder struct {
	Id        string   `json:"id"`
	Car       Car      `json:"car"`
	Customer  Customer `json:"customer"`
	FromDate  string   `json:"from_date"`
	ToDate    string   `json:"to_date"`
	Status    string   `json:"status"`
	Paid      bool     `json:"payment_status"`
	Amount     float32    `json:"amount"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
	
}


type CreateOrder struct {
	CarId      string `json:"car_id"`
	CustomerId string `json:"customer_id"`
	FromDate   string `json:"from_date"`
	ToDate     string `json:"to_date"`
	Status     string `json:"status"`
	Paid       bool   `json:"payment_status"`
	Amount     float32    `json:"amount"`
}

type GetAllOrders struct {
	Orders []GetOrder `json:"orders"`
	Count  int        `json:"count"`
}

type GetAllOrdersRequest struct {
    Search string `json:"search"`
	Page uint64 `json:"page"`
	Limit uint64 `json:"limit"`
}

type GetAllOrdersResponse struct {
	Orders []GetAllOrders `json:"orders"`
	Count int16 `json:"count"`
}

package main

import (
	"clone/rent_car_us/config"
	"clone/rent_car_us/controller"
	"clone/rent_car_us/storage/postgres"
	"fmt"
	"net/http"
)
func main() {
	cfg := config.Load()
	store, err := postgres.New(cfg)
	if err != nil {
		fmt.Println("error while connecting db, err: ", err)
		return
	}
	defer store.CloseDB()

	con := controller.NewController(store)

	http.HandleFunc("/car", con.Car)
	http.HandleFunc("/customer", con.Customer)
	http.HandleFunc("/order", con.Order)

	fmt.Println("programm is running on localhost:8008...")
	http.ListenAndServe(":8080", nil)

}

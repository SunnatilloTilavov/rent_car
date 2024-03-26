package main

import (
	"context"
	"fmt"
	"clone/rent_car_us/api"
	"clone/rent_car_us/config"
	"clone/rent_car_us/storage/postgres"
)

func main() {
	cfg := config.Load()
	store, err := postgres.New(context.Background(), cfg)
	if err != nil {
		fmt.Println("error while connecting db, err: ", err)
		return
	}
	defer store.CloseDB()

	c := api.New(store)

	fmt.Println("programm is running on localhost:8080...")
	c.Run(":8080")
}


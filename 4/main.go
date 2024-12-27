package main

import (
	"four/modules"
	"four/transport"

	"fmt"
	"net/http"
)

func setupRoutes(db *modules.DB) {
	handler := transport.NewBaseHandlerWithTableCustomers(db)
	http.HandleFunc("/api/v1/create-customer", handler.CreateCustomer)
	http.HandleFunc("/api/vi/authentication", handler.AuthCustomer)
}

func main() {
	db, err := modules.NewDB()
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting server")
	setupRoutes(db)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v", err)
	}
}

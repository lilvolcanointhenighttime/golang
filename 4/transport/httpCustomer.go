package transport

import (
	"four/modules"

	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
)

func (h *baseHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create Endpoint Hit")
	var user modules.Customer
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	_, err = mail.ParseAddress(user.Email)
	if err != nil {
		fmt.Println("1111111")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.Password, err = modules.HashPassword(user.Password)
	if err != nil {
		fmt.Println("2222")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.db.CheckCustomer(user)
	if err != nil {
		fmt.Println("3333333")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err = h.db.CreateCustomer(user)
	if err != nil {
		fmt.Println("44444")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *baseHandler) AuthCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("auth Endpoint Hit")
	var auth modules.AuthRequest
	err := json.NewDecoder(r.Body).Decode(&auth)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	auth.Password, err = modules.HashPassword(auth.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = h.db.GetCustomer(auth.Email, auth.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

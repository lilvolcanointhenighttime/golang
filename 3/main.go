package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

type Task3Schema struct {
	Text string `json:"text" validate:"required"`
}

func task1Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("task1 Endpoint Hit")
	if r.Method != http.MethodGet {
		http.Error(w, "only get method is supported", http.StatusBadRequest)
		return
	}
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")

	if name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}
	if age == "" {
		http.Error(w, "age is required", http.StatusBadRequest)
		return
	}

	ageInt, err := strconv.Atoi(age)
	if err != nil || ageInt <= 0 {
		http.Error(w, "age must be a positive integer", http.StatusBadRequest)
		return
	}

	response := Response{
		Message: fmt.Sprintf("Меня зовут %s, и мне %d лет.", name, ageInt),
		Status:  "success",
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func task2AddHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("task2Add Endpoint Hit")
	if r.Method != http.MethodGet {
		http.Error(w, "only get method is supported", http.StatusBadRequest)
		return
	}
	a := r.URL.Query().Get("a")
	b := r.URL.Query().Get("b")

	if a == "" {
		http.Error(w, "a is required", http.StatusBadRequest)
		return
	}
	if b == "" {
		http.Error(w, "b is required", http.StatusBadRequest)
		return
	}

	aFloat, err := strconv.ParseFloat(a, 64)
	if err != nil {
		http.Error(w, "a must be a number", http.StatusBadRequest)
		return
	}

	bFloat, err := strconv.ParseFloat(b, 64)
	if err != nil {
		http.Error(w, "b must be a number", http.StatusBadRequest)
		return
	}

	response := Response{
		Message: fmt.Sprintf("%f", (aFloat + bFloat)),
		Status:  "success",
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func task2SubHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("task2Sub Endpoint Hit")
	if r.Method != http.MethodGet {
		http.Error(w, "only get method is supported", http.StatusBadRequest)
		return
	}
	a := r.URL.Query().Get("a")
	b := r.URL.Query().Get("b")

	if a == "" {
		http.Error(w, "a is required", http.StatusBadRequest)
		return
	}
	if b == "" {
		http.Error(w, "b is required", http.StatusBadRequest)
		return
	}

	aFloat, err := strconv.ParseFloat(a, 64)
	if err != nil {
		http.Error(w, "a must be a number", http.StatusBadRequest)
		return
	}

	bFloat, err := strconv.ParseFloat(b, 64)
	if err != nil {
		http.Error(w, "b must be a number", http.StatusBadRequest)
		return
	}

	response := Response{
		Message: fmt.Sprintf("%f", (aFloat - bFloat)),
		Status:  "success",
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func task2MulHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("task2Mul Endpoint Hit")
	if r.Method != http.MethodGet {
		http.Error(w, "only get method is supported", http.StatusBadRequest)
		return
	}
	a := r.URL.Query().Get("a")
	b := r.URL.Query().Get("b")

	if a == "" {
		http.Error(w, "a is required", http.StatusBadRequest)
		return
	}
	if b == "" {
		http.Error(w, "b is required", http.StatusBadRequest)
		return
	}

	aFloat, err := strconv.ParseFloat(a, 64)
	if err != nil {
		http.Error(w, "a must be a number", http.StatusBadRequest)
		return
	}

	bFloat, err := strconv.ParseFloat(b, 64)
	if err != nil {
		http.Error(w, "b must be a number", http.StatusBadRequest)
		return
	}

	response := Response{
		Message: fmt.Sprintf("%f", (aFloat * bFloat)),
		Status:  "success",
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func task2DivHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("task2Div Endpoint Hit")
	if r.Method != http.MethodGet {
		http.Error(w, "only get method is supported", http.StatusBadRequest)
		return
	}
	a := r.URL.Query().Get("a")
	b := r.URL.Query().Get("b")

	if a == "" {
		http.Error(w, "a is required", http.StatusBadRequest)
		return
	}
	if b == "" {
		http.Error(w, "b is required", http.StatusBadRequest)
		return
	}

	aFloat, err := strconv.ParseFloat(a, 64)
	if err != nil {
		http.Error(w, "a must be a number", http.StatusBadRequest)
		return
	}

	bFloat, err := strconv.ParseFloat(b, 64)
	if err != nil {
		http.Error(w, "b must be a number", http.StatusBadRequest)
		return
	}

	response := Response{
		Message: fmt.Sprintf("%f", (aFloat / bFloat)),
		Status:  "success",
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func countCharacters(s string) (map[string]int, error) {
	counts := make(map[string]int)

	for _, char := range s {
		counts[string(char)]++
	}

	return counts, nil
}

func task3Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("task3 Endpoint Hit")

	if r.Method != http.MethodPost {
		http.Error(w, "only post method is supported", http.StatusBadRequest)
		return
	}

	var body Task3Schema
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	charCounts, err := countCharacters(body.Text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonData, err := json.Marshal(charCounts)
	if err != nil {
		fmt.Println("Ошибка маршалинга:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := Response{
		Message: string(jsonData),
		Status:  "success",
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func setupRoutes() {
	http.HandleFunc("/api/v1/task1", task1Handler)
	http.HandleFunc("/api/v1/task2/add", task2AddHandler)
	http.HandleFunc("/api/v1/task2/sub", task2SubHandler)
	http.HandleFunc("/api/v1/task2/mul", task2MulHandler)
	http.HandleFunc("/api/v1/task2/div", task2DivHandler)
	http.HandleFunc("/api/v1/task3", task3Handler)
}

func main() {
	fmt.Println("Starting server")
	setupRoutes()
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v", err)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	Name string `json:"name"`
}

type UserResponse struct {
	Massage string `json:"massage"`
}

func main() {
	fmt.Println("Started listening port: 8081")
	http.HandleFunc("/echo", helloBody)
	http.HandleFunc("/greet", greet)
	http.ListenAndServe(":8081", nil)
}

func helloBody(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Reading body error", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Got: %s", body)
}

func greet(w http.ResponseWriter, r *http.Request) {
	var (
		user User
		resp UserResponse
	)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Reading body error", http.StatusBadRequest)
		return
	}
	resp = UserResponse{Massage: "Hello," + user.Name}
	json.NewEncoder(w).Encode(resp)
}

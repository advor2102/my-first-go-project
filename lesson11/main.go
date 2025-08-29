package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/echo", helloBody)
	fmt.Println("Started listening port: 8081")
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

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Message struct {
	Text string `json:"text"`
}

func helloAPI(w http.ResponseWriter, r *http.Request) {
	message := Message{Text: "Hello, Go Microservice!"}
	json.NewEncoder(w).Encode(message)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/hello", helloAPI).Methods("GET")
	fmt.Println("Microservice running on :8080")
	http.ListenAndServe(":8080", r)
}

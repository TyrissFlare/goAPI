package main

import (
	"encoding/json"
	"net/http"
)

// Message represent a simple JSON message
type Message struct {
	Greeting string `json:"greeting"`
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	message := Message{Greeting: "Hello, World!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func main() {
	http.HandleFunc("/welcome", greetHandler)
	http.ListenAndServe(":8011", nil)
}
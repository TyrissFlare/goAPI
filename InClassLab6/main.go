package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

// GreetingResponse represents the response structure for the greeting endpoint
type GreetingResponse struct {
	Message string `json:"message"`
}

// Datastore represents a simple in-memory key-value store
type Datastore struct {
	data map[string]string
	mu   sync.RWMutex
}

// NewDatastore initializes a new Datastore
func NewDatastore() *Datastore {
	return &Datastore{
		data: make(map[string]string),
	}
}

// Set stores a key-value pair in the datastore
func (d *Datastore) Set(key, value string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.data[key] = value
}

// Get retrieves the value associated with the given key from the datastore
func (d *Datastore) Get(key string) (string, bool) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	value, ok := d.data[key]
	return value, ok
}

func main() {
	// Initialize datastore
	store := NewDatastore()

	// Define REST API endpoints
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		// Return a greeting message
		response := GreetingResponse{Message: "Hello, World!"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	http.HandleFunc("/store", func(w http.ResponseWriter, r *http.Request) {
		// Store a key-value pair in the datastore
		key := r.URL.Query().Get("key")
		value := r.URL.Query().Get("value")
		if key == "" || value == "" {
			http.Error(w, "Key and value must be provided", http.StatusBadRequest)
			return
		}
		store.Set(key, value)
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/retrieve", func(w http.ResponseWriter, r *http.Request) {
		// Retrieve a value associated with the given key from the datastore
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "Key must be provided", http.StatusBadRequest)
			return
		}
		value, ok := store.Get(key)
		if !ok {
			http.Error(w, "Key not found", http.StatusNotFound)
			return
		}
		response := GreetingResponse{Message: value}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	// Start HTTP server
	log.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
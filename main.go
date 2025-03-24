package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "SUCCESS"}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/health", healthHandler)
	
	port := ":8081"
	log.Printf("Starting server on port %s...", port)
	
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
} 
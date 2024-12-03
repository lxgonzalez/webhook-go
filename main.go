package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type WebhookPayload struct {
	Event   string `json:"event"`
	Message string `json:"message"`
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Parse the JSON payload
	var payload WebhookPayload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	// Log the payload for debugging
	log.Printf("Received webhook: Event=%s, Message=%s", payload.Event, payload.Message)

	// Send a response back
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Webhook received: %s", payload.Event)
}

func main() {
	http.HandleFunc("/webhook", webhookHandler)

	port := "8080"
	log.Printf("Starting server on port %s...", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// Write a simple response to the client
	w.WriteHeader(http.StatusOK)
	// Return a JSON response with a status message
	w.Header().Set("Content-Type", "application/json")
	// Use a struct to define the JSON response format
	// and write the response to the client
	fmt.Fprintf(w, `{"status":"available", "environment": "%s", "version": "%s"}`, app.config.env, version)
}

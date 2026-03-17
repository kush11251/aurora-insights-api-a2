package controllers

import (
	"encoding/json"
	"net/http"

	"aurora-insights-api/src/models"
)

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	// Fetch user data from database or other data source
	// For demonstration purposes, return a hardcoded user
	user := models.User{ID: "1", Username: "john", Email: "john@example.com"}
	json.NewEncoder(w).Encode(user)
}
package handlers

import (
	"idp/internal/services"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if username == "" || email == "" || password == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	err := services.CreateUser(username, email, password)
	if err != nil {
		http.Error(w, "Failed to create user:"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered successfully"))
}

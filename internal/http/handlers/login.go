package handlers

import (
	"html/template"
	"net/http"

	"IdP/internal/session"
	"IdP/internal/store"
)

var loginTpl = template.Must(template.ParseFiles("web/templates/login.html"))

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		redirect := r.URL.Query().Get("redirect")
		loginTpl.Execute(w, map[string]string{"Redirect": redirect})

	case http.MethodPost:
		username := r.FormValue("username")
		password := r.FormValue("password")
		redirect := r.FormValue("redirect")

		if store.Users[username] != password {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}

		session.SetUser(w, username)
		http.Redirect(w, r, redirect, http.StatusFound)
	}
}

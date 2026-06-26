package handlers

import (
	"log"
	"net/http"
	"net/url"

	"IdP/internal/auth/tokens"
)

func AuthorizeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Only allow GET for now
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		q := r.URL.Query()

		clientID := q.Get("client_id")
		redirectURI := q.Get("redirect_uri")
		responseType := q.Get("response_type")
		state := q.Get("state")

		if clientID == "" || redirectURI == "" || responseType == "" {
			http.Error(w, "missing required parameters", http.StatusBadRequest)
			return
		}

		if responseType != "code" {
			http.Error(w, "unsupported response_type", http.StatusBadRequest)
			return
		}

		// TODO: validate client_id and redirect_uri against registered clients
		// For now, we trust them.

		// Fake user for MVP
		userID := "user-123"

		ac, err := tokens.GenerateAuthCode(clientID, userID, redirectURI)
		if err != nil {
			log.Println("failed to generate auth code:", err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		// Build redirect URL: redirect_uri?code=...&state=...
		redirect, err := url.Parse(redirectURI)
		if err != nil {
			http.Error(w, "invalid redirect_uri", http.StatusBadRequest)
			return
		}

		params := redirect.Query()
		params.Set("code", ac.Code)
		if state != "" {
			params.Set("state", state)
		}
		redirect.RawQuery = params.Encode()

		http.Redirect(w, r, redirect.String(), http.StatusFound)
	}
}

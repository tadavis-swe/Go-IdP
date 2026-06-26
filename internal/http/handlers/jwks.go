package handlers

import (
	"encoding/json"
	"net/http"

	"IdP/internal/auth/tokens"
)

type JWKSResponse struct {
	Keys []tokens.JWK `json:"keys"`
}

func JWKSHandler(jwk tokens.JWK) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(JWKSResponse{Keys: []tokens.JWK{jwk}})
	}
}

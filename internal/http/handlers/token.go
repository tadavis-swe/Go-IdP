package handlers

import (
	"crypto/rsa"
	"encoding/json"
	"net/http"

	"IdP/internal/auth/tokens"
)

type TokenHandlerDeps struct {
	PrivateKey *rsa.PrivateKey
	Issuer     string
}

func TokenHandler(deps TokenHandlerDeps) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, "invalid form", http.StatusBadRequest)
			return
		}

		grantType := r.Form.Get("grant_type")
		code := r.Form.Get("code")
		redirectURI := r.Form.Get("redirect_uri")
		clientID := r.Form.Get("client_id")

		if grantType != "authorization_code" {
			http.Error(w, "unsupported grant_type", http.StatusBadRequest)
			return
		}

		ac, ok := tokens.GetAuthCode(code)
		if !ok {
			http.Error(w, "invalid code", http.StatusBadRequest)
			return
		}

		if ac.ClientID != clientID || ac.RedirectURI != redirectURI {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		// Consume the code
		tokens.DeleteAuthCode(code)

		idToken, err := tokens.GenerateIDToken(deps.PrivateKey, deps.Issuer, clientID, ac.UserID)
		if err != nil {
			http.Error(w, "failed to generate id token", http.StatusInternalServerError)
			return
		}

		// Access token can be opaque for now
		accessToken := "access-" + code

		resp := map[string]interface{}{
			"access_token": accessToken,
			"id_token":     idToken,
			"token_type":   "Bearer",
			"expires_in":   3600,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

package main

import (
	"log"
	"net/http"

	"IdP/internal/auth/tokens"
	"IdP/internal/http/handlers"
)

func main() {
	issuer := "http://localhost:8080"

	// Generate signing key
	privKey, err := tokens.GenerateRSAKey()
	if err != nil {
		log.Fatal(err)
	}

	jwk := tokens.PublicJWKFromKey(privKey)

	http.HandleFunc("/.well-known/openid-configuration", handlers.DiscoveryHandler(issuer))
	http.HandleFunc("/jwks.json", handlers.JWKSHandler(jwk))
	http.HandleFunc("/authorize", handlers.AuthorizeHandler())
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/token", handlers.TokenHandler(handlers.TokenHandlerDeps{
		PrivateKey: privKey,
		Issuer:     issuer,
	}))

	log.Println("IdP running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

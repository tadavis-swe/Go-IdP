package tokens

import (
	"crypto/rsa"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateIDToken(priv *rsa.PrivateKey, issuer, clientID, userID string) (string, error) {
	claims := jwt.MapClaims{
		"iss": issuer,
		"sub": userID,
		"aud": clientID,
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(priv)
}

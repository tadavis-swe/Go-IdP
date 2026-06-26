package tokens

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"math/big"
)

type JWK struct {
	Kty string `json:"kty"`
	Kid string `json:"kid"`
	Alg string `json:"alg"`
	Use string `json:"use"`
	N   string `json:"n"`
	E   string `json:"e"`
}

func GenerateRSAKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}

func PublicJWKFromKey(key *rsa.PrivateKey) JWK {
	pub := key.PublicKey

	n := base64.RawURLEncoding.EncodeToString(pub.N.Bytes())
	e := base64.RawURLEncoding.EncodeToString(big.NewInt(int64(pub.E)).Bytes())

	return JWK{
		Kty: "RSA",
		Kid: "primary",
		Alg: "RS256",
		Use: "sig",
		N:   n,
		E:   e,
	}
}

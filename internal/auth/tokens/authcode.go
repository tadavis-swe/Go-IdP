package tokens

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
)

type AuthCode struct {
	Code        string
	ClientID    string
	UserID      string
	RedirectURI string
}

var (
	authCodes   = make(map[string]AuthCode)
	authCodesMu sync.Mutex
)

func GenerateAuthCode(clientID, userID, redirectURI string) (AuthCode, error) {
	buf := make([]byte, 32)
	if _, err := rand.Read(buf); err != nil {
		return AuthCode{}, err
	}

	code := hex.EncodeToString(buf)

	ac := AuthCode{
		Code:        code,
		ClientID:    clientID,
		UserID:      userID,
		RedirectURI: redirectURI,
	}

	authCodesMu.Lock()
	authCodes[code] = ac
	authCodesMu.Unlock()

	return ac, nil
}

func GetAuthCode(code string) (AuthCode, bool) {
	authCodesMu.Lock()
	defer authCodesMu.Unlock()

	ac, ok := authCodes[code]
	return ac, ok
}

func DeleteAuthCode(code string) {
	authCodesMu.Lock()
	delete(authCodes, code)
	authCodesMu.Unlock()
}

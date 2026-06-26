package session

import (
	"net/http"
)

const CookieName = "session_user"

func SetUser(w http.ResponseWriter, username string) {
	http.SetCookie(w, &http.Cookie{
		Name:  CookieName,
		Value: username,
		Path:  "/",
	})
}

func GetUser(r *http.Request) (string, bool) {
	c, err := r.Cookie(CookieName)
	if err != nil {
		return "", false
	}
	return c.Value, true
}

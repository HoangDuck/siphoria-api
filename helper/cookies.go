package helper

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"
)

func GenerateStateOauthCookie(responseWriter http.ResponseWriter) string {
	var expiration = time.Now().Add(2 * time.Minute)
	randomCode := make([]byte, 16)
	_, err := rand.Read(randomCode)
	if err != nil {
		return ""
	}
	state := base64.URLEncoding.EncodeToString(randomCode)
	cookie := http.Cookie{
		Name:     "oauthstate",
		Value:    state,
		Expires:  expiration,
		HttpOnly: true,
	}
	http.SetCookie(responseWriter, &cookie)
	return state
}

func GenerateAccessTokenOauthCookie(responseWriter http.ResponseWriter, token string) {
	var expiration = time.Now().Add(2 * time.Minute)
	cookie := http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  expiration,
		HttpOnly: true,
	}
	http.SetCookie(responseWriter, &cookie)
}

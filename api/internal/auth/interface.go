package auth

import (
	"net/http"

	"github.com/Stei-ITstudents/go-auth/api/internal/db"
)

type Authenticator interface {
	Login(userStore db.UserStore) http.HandlerFunc
	Logout(w http.ResponseWriter, r *http.Request)
	Middleware(next http.HandlerFunc, sessionID string) http.HandlerFunc
}

package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Stei-ITstudents/go-auth/api/internal/db"
	"github.com/gorilla/sessions"
)

type SessionAuthenticator struct {
	CookieStore *sessions.CookieStore
}

func NewSessionAuthenticator(secret string) *SessionAuthenticator {
	return &SessionAuthenticator{
		CookieStore: sessions.NewCookieStore([]byte(secret)),
	}
}

func (sa *SessionAuthenticator) Login(userStore db.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			username := r.FormValue("username")
			password := r.FormValue("password")

			if Authenticate(userStore, username, password) {
				session, _ := sa.CookieStore.Get(r, "user_session")
				session.Values["authenticated"] = true
				session.Values["user"] = username
				session.Save(r, w)
				http.Redirect(w, r, "/welcome", http.StatusFound)
			} else {
				http.Redirect(w, r, "/login", http.StatusFound)
			}
		} else {
			fmt.Fprintf(w, `<html><body><form method="POST" action="/login">
				Username: <input type="text" name="username"><br>
				Password: <input type="password" name="password"><br>
				<input type="submit" value="Login">
			</form></body></html>`)
		}
	}
}

func (sa *SessionAuthenticator) Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := sa.CookieStore.Get(r, "user_session")
	session.Values["authenticated"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusFound)

}

func (sa *SessionAuthenticator) Middleware(next http.HandlerFunc, sessionID string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := sa.CookieStore.Get(r, sessionID)
		auth, ok := session.Values["authenticated"].(bool)
		if !ok || !auth {
			log.Print()
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		next.ServeHTTP(w, r)
	}
}

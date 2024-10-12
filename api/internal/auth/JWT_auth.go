package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Stei-ITstudents/go-auth/api/internal/db"
	"github.com/dgrijalva/jwt-go"
)

type JWTAuthenticator struct {
	Secret string
}

func NewJWTAuthenticator(secret string) *JWTAuthenticator {
	return &JWTAuthenticator{
		Secret: secret,
	}
}

func (ja *JWTAuthenticator) Login(userStore db.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			username := r.FormValue("username")
			password := r.FormValue("password")

			if Authenticate(userStore, username, password) {
				token, err := ja.GenerateToken(username)
				if err != nil {
					http.Error(w, "Could not generate token", http.StatusInternalServerError)
					return
				}
				http.SetCookie(w, &http.Cookie{
					Name:    "Authorization",
					Value:   token,
					Expires: time.Now().Add(24 * time.Hour),
				})
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

func (ja *JWTAuthenticator) GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(ja.Secret))
}

func (ja *JWTAuthenticator) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "Authorization",
		Value:  "",
		MaxAge: -1,
	})

	http.Redirect(w, r, "/login", http.StatusFound)
}

func (ja *JWTAuthenticator) Middleware(next http.HandlerFunc, sessionID string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Authorization")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		tokenString := cookie.Value

		_, err = ja.VerifyToken(tokenString)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func (ja *JWTAuthenticator) VerifyToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(ja.Secret), nil
	})
}

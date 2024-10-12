package auth

import (
	"fmt"
	"net/http"

	"github.com/Stei-ITstudents/go-auth/api/internal/db"
	"golang.org/x/crypto/bcrypt"
)

func Signup(userStore db.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			username := r.FormValue("username")
			password := r.FormValue("password")

			err := userStore.Add(username, password)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			} else {
				_, _ = fmt.Fprintf(w, "User %s added succesifull!", username)
			}

		} else {
			fmt.Fprintf(w, `<html><body><form method="POST" action="/signup">
				Username: <input type="text" name="username"><br>
				Password: <input type="password" name="password"><br>
				<input type="submit" value="Add User">
			</form></body></html>`)
		}
	}
}

func Authenticate(userStore db.UserStore, username, password string) bool {
	hashedPassword, err := userStore.Get(username)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}

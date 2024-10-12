package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Stei-ITstudents/go-auth/api/internal/auth"
	"github.com/Stei-ITstudents/go-auth/api/internal/db"
	"github.com/caarlos0/env/v9"
)

type Config struct {
	Port      int    `env:"PORT" envDefault:"8080"`
	SecretKey string `env:"SECRET_KEY" envDefault:"my-Secret-Key"`
	MySqlDns  string `env:"DB_DNS" envDefault:"IMPLEMENT_IT!!!!"`
}

func GetConfigs() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

type Application struct {
	session_id string
}

func NewApplication(session_id string) *Application {
	return &Application{
		session_id: session_id,
	}
}

func (svc *Application) WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	// TODO solve getting username in handler fixing the below code
	// session, _ := auth.CookieStore.Get(r, svc.session_id)
	// user := session.Values["user"].(string)
	fmt.Fprint(w, "Welcome \n")
}

func main() {
	configs, err := GetConfigs()
	if err != nil {
		panic(fmt.Errorf("cannot get configs: %v", err))
	}

	userStore, err := db.NewInMemoryUserStore()
	if err != nil {
		log.Fatalf("Cannot initialize DB: %v", err)
	}
	userStore.Add("john", "password")
	// authenticator := auth.NewSessionAuthenticator(configs.SecretKey)
	// sau
	authenticator := auth.NewJWTAuthenticator(configs.SecretKey)

	sesion_id := "user_session"
	svc := NewApplication(sesion_id)

	http.HandleFunc("/welcome", authenticator.Middleware(svc.WelcomeHandler, sesion_id))
	http.HandleFunc("/login", authenticator.Login(userStore))
	http.HandleFunc("/logout", authenticator.Logout)
	http.HandleFunc("/signup", auth.Signup(userStore))
	log.Printf("Starting server on \033]8;;http://localhost:%d/login\033\\http://localhost:%d/login\033]8;;\033\\ ...\n", configs.Port, configs.Port)
	log.Fatal(http.ListenAndServe(strings.Join([]string{":", strconv.Itoa(configs.Port)}, " "), nil))
}

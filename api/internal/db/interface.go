package db

type UserStore interface {
	Get(username string) (string, error)
	Add(username string, password string) error
}

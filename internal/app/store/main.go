package store

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

type Store struct {
	UserRepository UserRepository
}

func New(dbString string, log *logrus.Logger) *Store {
	db, err := sql.Open("postgres", dbString)

	if err != nil {
		log.Panic("Error with open connection to DB", err)
	}

	if err := db.Ping(); err != nil {
		log.Panic("Error with pinging DB")
	}

	store := Store{}

	store.UserRepository = NewUserRepository(db, log)

	return &store
}

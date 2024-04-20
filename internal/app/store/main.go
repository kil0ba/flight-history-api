package store

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type Store struct {
	UserRepository  UserRepository
	PlaneRepository PlaneRepository
}

func New(ctx context.Context, dbString string, log *logrus.Logger) *Store {
	conn, err := pgx.Connect(ctx, dbString)

	if err != nil {
		log.Panic("Error with open connection to DB", err)
	}

	if err := conn.Ping(ctx); err != nil {
		log.Panic("Error with pinging DB")
	}

	store := Store{}

	store.UserRepository = NewUserRepository(conn, log)
	store.PlaneRepository = NewPlaneRepository(conn, log)

	return &store
}

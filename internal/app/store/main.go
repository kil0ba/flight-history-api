package store

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type Store struct {
	UserRepository    UserRepository
	PlaneRepository   PlaneRepository
	AirportRepository AirportRepository
	AirlineRepository AirlineRepository
}

func New(ctx context.Context, dbString string, log *logrus.Logger) *Store {
	dbpool, err := pgxpool.New(ctx, dbString)

	if err != nil {
		log.Panic("Error with open connection to DB", err)
	}

	if err := dbpool.Ping(ctx); err != nil {
		log.Panic("Error with pinging DB")
	}

	store := Store{}

	store.UserRepository = NewUserRepository(dbpool, log)
	store.PlaneRepository = NewPlaneRepository(dbpool, log)
	store.AirportRepository = NewAirportRepositry(dbpool, log)
	store.AirlineRepository = NewAirlineRepository(dbpool, log)

	return &store
}

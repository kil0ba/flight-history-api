package store

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	model "github.com/kil0ba/flight-history-api/internal/app/models"
	"github.com/sirupsen/logrus"
)

type AirlineRepository struct {
	db  *pgxpool.Pool
	log *logrus.Logger
}

func NewAirlineRepository(db *pgxpool.Pool, log *logrus.Logger) AirlineRepository {
	return AirlineRepository{
		db:  db,
		log: log,
	}
}

const airlineSearch = "Airline Search: "

func (ar *AirlineRepository) Search(ctx context.Context, query string, page int, count int) ([]model.Airline, error) {
	if page < 1 {
		page = 1
	}
	if count < 1 || count > 50 {
		count = searchCount
	}

	searchQuery := "%" + strings.ToLower(query) + "%"
	rows, err := ar.db.Query(ctx, `
	SELECT
	    airlines.*
	FROM
	    airlines
	JOIN
	(
	    SELECT
	        id,
	        LOWER(COALESCE(name,'') || COALESCE(alias,'') || COALESCE(iata ,'') || COALESCE(icao ,'') || COALESCE(callsign ,'') || COALESCE(country ,'')) AS concatenated
	    FROM
	        airlines
	) T ON T.id = airlines.id
	WHERE
	    T.concatenated LIKE $1
	LIMIT $2 OFFSET $3;
		`, searchQuery, count, (page-1)*count,
	)

	if err != nil {
		ar.log.WithError(err).Error(airlineSearch, "Failed to search")
		return nil, err
	}

	airlines := []model.Airline{}
	defer rows.Close()
	for rows.Next() {
		airline := model.Airline{}
		err := rows.Scan(
			&airline.ID,
			&airline.Name,
			&airline.Active,
			&airline.Alias,
			&airline.Country,
			&airline.IATA,
			&airline.ICAO,
			&airline.Callsign,
		)
		if err != nil {
			ar.log.WithError(err).Error(airlineSearch, "Failed to get row")
			return nil, err
		}
		airlines = append(airlines, airline)
	}

	return airlines, nil
}

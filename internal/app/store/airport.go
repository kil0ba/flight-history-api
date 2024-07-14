package store

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	model "github.com/kil0ba/flight-history-api/internal/app/models"
	"github.com/sirupsen/logrus"
)

type AirportRepository struct {
	db  *pgxpool.Pool
	log *logrus.Logger
}

func NewAirportRepositry(db *pgxpool.Pool, log *logrus.Logger) AirportRepository {
	return AirportRepository{
		db:  db,
		log: log,
	}
}

func (ar *AirportRepository) Create(ctx context.Context, airport *model.Airport) error {
	return ar.db.QueryRow(ctx, "INSERT INTO airports (name, code, city, country, latitude, longitude) VALUE ($1, $2, $3, $4, $5, $6)",
		airport.Name, airport.Code, airport.City, airport.Country, airport.Latitude, airport.Longitude,
	).Scan()
}

func (ar *AirportRepository) Get(ctx context.Context, id int) (*model.Airport, error) {
	airport := model.Airport{}
	err := ar.db.QueryRow(ctx, "SELECT (id, name, code, city, country, latitude, longitude) FROM airports  WHERE id = $1", id).Scan(
		&airport.ID, &airport.Name, &airport.Code, &airport.City, &airport.Country, &airport.Latitude, &airport.Longitude,
	)
	return &airport, err
}

const airportSearch = "Airport Search: "

func (ar *AirportRepository) Search(ctx context.Context, query string, page int, count int) ([]model.Airport, error) {
	if page < 1 {
		page = 1
	}

	airports := []model.Airport{}

	searchQuery := "%" + strings.ToLower(query) + "%"
	rows, err := ar.db.Query(ctx, `
	SELECT
	    airports.*
	FROM
	    airports
	JOIN
	(
	    SELECT
	        id,
	        LOWER(COALESCE(name,'') || COALESCE(code,'') || COALESCE(city ,'') || COALESCE(country ,'')) AS concatenated
	    FROM
	        airports
	) T ON T.id = airports.id
	WHERE
	    T.concatenated LIKE $1
	LIMIT $2 OFFSET $3;
		`, searchQuery, searchCount, (page-1)*searchCount,
	)

	if err != nil {
		ar.log.WithError(err).Error(airportSearch, "Failed to search aiports")
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		airport := model.Airport{}
		err := rows.Scan(&airport.ID, &airport.Name, &airport.Code, &airport.City, &airport.Country, &airport.Latitude, &airport.Longitude, &airport.Timezone)
		if err != nil {
			ar.log.WithError(err).Error(planeSearch, "Failed to get airport row")
			return nil, err
		}
		airports = append(airports, airport)
	}

	return airports, nil
}

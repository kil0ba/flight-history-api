package store

import (
	"context"

	"github.com/jackc/pgx/v5"
	model "github.com/kil0ba/flight-history-api/internal/app/models"
	"github.com/sirupsen/logrus"
)

type PlaneRepository struct {
	db  *pgx.Conn
	log *logrus.Logger
}

func NewPlaneRepository(db *pgx.Conn, log *logrus.Logger) PlaneRepository {
	return PlaneRepository{
		db:  db,
		log: log,
	}
}

func (pr *PlaneRepository) Create(ctx context.Context, p model.Plane) error {
	return pr.db.QueryRow(ctx, "INSERT INTO planes (name, iata_code, icao_code, manufacturer, country) VALUES ($1, $2, $3, $4, $5)",
		p.Name,
		p.IataCode,
		p.IcaoCode,
		p.Manufacturer,
		p.Country,
	).Scan()
}

func (pr *PlaneRepository) Get(ctx context.Context, id int) (*model.Plane, error) {
	plane := model.Plane{}
	if err := pr.db.QueryRow(
		ctx,
		"SELECT name, iata_code, icao_code, manufacturer, country from planes WHERE id = $1", id).Scan(
		&plane.Name, &plane.IataCode,
		&plane.IcaoCode, &plane.Manufacturer, &plane.Country); err != nil {
		return nil, err
	}
	return &plane, nil
}

const planeGetList = "Plane GetList "

func (pr *PlaneRepository) GetList(ctx context.Context, count int, page int) ([]model.Plane, error) {
	if page < 1 {
		page = 1
	}

	planes := []model.Plane{}

	rows, err := pr.db.Query(
		ctx,
		"SELECT name, iata_code, icao_code, manufacturer, country FROM planes ORDER BY manufacturer DESC LIMIT $1 OFFSET $2", count, (page-1)*count)

	if err != nil {
		pr.log.WithError(err).Error(planeGetList, "Failed to get planes")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		plane := model.Plane{}
		err := rows.Scan(&plane.Name, &plane.IataCode, &plane.IcaoCode, &plane.Manufacturer, &plane.Country)
		if err != nil {
			pr.log.WithError(err).Error(planeGetList, "Failed to get plane row")
			return nil, err
		}
		planes = append(planes, plane)
	}

	return planes, nil
}

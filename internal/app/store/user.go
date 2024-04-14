package store

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	model "github.com/kil0ba/flight-history-api/internal/app/models"
	"github.com/sirupsen/logrus"
)

type UserRepository struct {
	db  *pgx.Conn
	log *logrus.Logger
}

func NewUserRepository(db *pgx.Conn, log *logrus.Logger) UserRepository {
	return UserRepository{
		db:  db,
		log: log,
	}
}

func (r *UserRepository) Create(ctx context.Context, u model.User) error {
	return r.db.QueryRow(ctx,
		"INSERT INTO users (uuid, email, encrypted_password, login) VALUES ($1, $2, $3, $4) RETURNING email",
		uuid.New(),
		u.Email,
		u.Password,
		u.Login,
	).Scan(&u.Email)
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	user := model.User{}
	if err := r.db.QueryRow(ctx,
		"SELECT uuid, email, encrypted_password, login FROM users WHERE email = $1",
		email,
	).Scan(&user.Uuid, &user.Email, &user.EncryptedPassword, &user.Login); err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("record not found")
		}
		return nil, fiber.ErrInternalServerError
	}

	return &user, nil
}

func (r *UserRepository) FindByLogin(ctx context.Context, login string) (*model.User, error) {
	user := model.User{}
	if err := r.db.QueryRow(ctx,
		"SELECT uuid, email, encrypted_password, login FROM users WHERE login = $1",
		login,
	).Scan(&user.Uuid, &user.Email, &user.EncryptedPassword, &user.Login); err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("record not found")
		}
		return nil, fiber.ErrInternalServerError
	}

	return &user, nil
}

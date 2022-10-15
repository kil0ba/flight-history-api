package store

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	model "github.com/kil0ba/flight-history-api/internal/app/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(u model.User) error {
	return r.db.QueryRow(
		"INSERT INTO users (uuid, email, encrypted_password, login) VALUES ($1, $2, $3, $4) RETURNING email",
		uuid.New(),
		u.Email,
		u.Password,
		u.Login,
		).Scan(&u.Uuid)
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	user := model.User{}
	if err := r.db.QueryRow(
		"SELECT uuid, email, encrypted_password, login FROM users WHERE email = $1",
		email,
		).Scan(&user.Uuid, &user.Email, &user.EncryptedPassword, &user.Login); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("record not found")
		}
		return nil, fiber.ErrInternalServerError
	}

	return &user, nil
}

func (r *UserRepository) FindByLogin(login string)  (*model.User, error) {
	user := model.User{}
	if err := r.db.QueryRow(
		"SELECT uuid, email, encrypted_password, login FROM users WHERE login = $1",
		login,
		).Scan(&user.Uuid, &user.Email, &user.EncryptedPassword, &user.Login); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("record not found")
		}
		return nil, fiber.ErrInternalServerError
		}

		return &user, nil
}
package repository

import (
	"database/sql"

	"github.com/muhangga/internal/entity"
)

type AuthRepository interface {
	Save(user entity.User) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
}

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *authRepository {
	return &authRepository{db}
}

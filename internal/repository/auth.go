package repository

import (
	"database/sql"

	"github.com/muhangga/internal/entity"
)

type AuthRepository interface {
	Save(user entity.User) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	FindByID(id int64) (entity.User, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db}
}

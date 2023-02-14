package repository

import (
	"database/sql"

	"github.com/muhangga/internal/entity"
)

type UserRepository interface {
	GetAllUser() ([]entity.User, error)
	UpdateUser(e entity.User) (entity.User, error)
	FindByID(id int64) (entity.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

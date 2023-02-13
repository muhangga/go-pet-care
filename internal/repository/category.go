package repository

import (
	"database/sql"

	"github.com/muhangga/internal/entity"
)

type CategoryRepository interface {
	Save(category entity.Category) (entity.Category, error)
	FindAll() ([]entity.Category, error)
	Update(category entity.Category) (entity.Category, error)
	Delete(id int64) error
}

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *categoryRepository {
	return &categoryRepository{db}
}

package repository

import (
	"database/sql"

	"github.com/muhangga/internal/entity"
)

type SpecialtiesRepository interface {
	FindAllSpecialties() ([]entity.Specialties, error)
	Save(s entity.Specialties) (entity.Specialties, error)
}

type specialtiesRepository struct {
	db *sql.DB
}

func NewSpecialtiesRepository(db *sql.DB) *specialtiesRepository {
	return &specialtiesRepository{db}
}

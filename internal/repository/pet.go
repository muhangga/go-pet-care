package repository

import (
	"database/sql"

	"github.com/muhangga/internal/entity"
)

type PetRepository interface {
	Save(r entity.Pet) (entity.Pet, error)
}

type petRepository struct {
	db *sql.DB
}

func NewPetRepository(db *sql.DB) *petRepository {
	return &petRepository{db}
}

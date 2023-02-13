package entity

import "time"

type Specialties struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Popular   bool      `json:"popular"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SpecialtiesRequest struct {
	ID      int    `json:"id"`
	Name    string `json:"name" validate:"required"`
	Popular bool   `json:"popular" validate:"required"`
}

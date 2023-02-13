package entity

import "time"

type Category struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CategoryRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name" validate:"required"`
	Icon string `json:"icon" validate:"required"`
}

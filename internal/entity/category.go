package entity

import "time"

type Category struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Icon      string    `json:"icon" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;not null"`
}

type CategoryRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name" validate:"required"`
	Icon string `json:"icon" validate:"required"`
}

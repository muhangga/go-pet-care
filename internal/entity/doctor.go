package entity

import "time"

type Doctor struct {
	ID            int64     `json:"id" gorm:"primary_key"`
	SpecialtiesID int64     `json:"specialties_id" gorm:"type:bigint;not null"`
	Name          string    `json:"name" gorm:"type:varchar(255);not null"`
	Rating        float64   `json:"rating" gorm:"type:float;not null"`
	Reviews       int64     `json:"reviews" gorm:"type:bigint;not null"`
	Experience    int64     `json:"experience" gorm:"type:bigint;not null"`
	Price         int64     `json:"price" gorm:"type:bigint;not null"`
	Address       string    `json:"address" gorm:"type:varchar(255);not null"`
	CreatedAt     time.Time `json:"created_at" gorm:"type:timestamp;not null"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"type:timestamp;not null"`
}

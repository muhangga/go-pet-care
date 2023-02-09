package entity

import "time"

type Pet struct {
	ID           int64     `json:"id" gorm:"primary_key"`
	UserID       int64     `json:"user_id" gorm:"type:bigint;not null"`
	AdditionalID int64     `json:"additional_id" gorm:"type:bigint;not null"`
	PetName      string    `json:"pet_name" gorm:"type:varchar(255);not null"`
	Gender       string    `json:"gender" gorm:"type:varchar(255);not null"`
	DateOfBirth  time.Time `json:"date_of_birth" gorm:"type:timestamp;not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"type:timestamp;not null"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"type:timestamp;not null"`
}

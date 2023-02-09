package entity

import "time"

type User struct {
	ID          int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	FullName    string    `json:"full_name" gorm:"type:varchar(255);not null"`
	Email       string    `json:"email" gorm:"type:varchar(255);not null"`
	Password    string    `json:"password" gorm:"type:varchar(255);not null"`
	PhoneNumber string    `json:"phone_number" gorm:"type:varchar(255);null"`
	Gender      string    `json:"gender" gorm:"type:varchar(255);not null"`
	Avatar      string    `json:"avatar" gorm:"type:varchar(255);null"`
	AboutMe     string    `json:"about_me" gorm:"type:varchar(255);null"`
	Role        string    `json:"role" gorm:"type:varchar(255);not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"type:timestamp;not null"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"type:timestamp;not null"`
}

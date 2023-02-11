package entity

import (
	"time"
)

type User struct {
	ID          int64     `json:"id"  database:"id"`
	FullName    string    `json:"full_name" database:"full_name"`
	Email       string    `json:"email" database:"email"`
	Password    string    `json:"password" database:"password"`
	PhoneNumber string    `json:"phone_number" database:"phone_number"`
	Gender      string    `json:"gender" database:"gender"`
	Avatar      string    `json:"avatar" database:"avatar"`
	AboutMe     string    `json:"about_me" database:"about_me"`
	Role        string    `json:"role" database:"role"`
	CreatedAt   time.Time `json:"created_at" database:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" database:"updated_at"`
}

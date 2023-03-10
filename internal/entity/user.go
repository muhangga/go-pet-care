package entity

import (
	"time"
)

type User struct {
	ID          int64     `json:"id" `
	FullName    string    `json:"full_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	PhoneNumber *string   `json:"phone_number"`
	Gender      *string   `json:"gender"`
	Avatar      string    `json:"avatar"`
	AboutMe     *string   `json:"about_me"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserUpdateRequest struct {
	ID        int64     `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Gender    *string   `json:"gender"`
	Avatar    string    `json:"avatar"`
	AboutMe   *string   `json:"about_me"`
	UpdatedAt time.Time `json:"updated_at"`
}

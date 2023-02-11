package utils

import (
	"time"

	"github.com/muhangga/internal/entity"
)

type UserFormatter struct {
	ID          int64     `json:"id"`
	FullName    string    `json:"full_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	PhoneNumber string    `json:"phone_number"`
	Gender      string    `json:"gender"`
	Avatar      string    `json:"avatar"`
	AboutMe     string    `json:"about_me"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at "`
}

func FormatUser(user entity.User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:          user.ID,
		FullName:    user.FullName,
		Email:       user.Email,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
		Gender:      user.Gender,
		Avatar:      user.Avatar,
		AboutMe:     user.AboutMe,
		Role:        user.Role,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}

	return formatter
}

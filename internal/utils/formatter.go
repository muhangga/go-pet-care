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

type LoginSuccessFormatter struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func FormatUser(user entity.User) UserFormatter {
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

func FormatLoginSuccess(user entity.User, token string) LoginSuccessFormatter {
	formatter := LoginSuccessFormatter{
		ID:    user.ID,
		Email: user.Email,
		Token: token,
	}

	return formatter
}

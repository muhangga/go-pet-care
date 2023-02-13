package entity

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role"`
}

type CheckEmailExist struct {
	Email string `json:"email" binding:"required,email"`
}

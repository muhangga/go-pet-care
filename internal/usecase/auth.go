package usecase

import (
	"github.com/muhangga/internal/entity"
	"github.com/muhangga/internal/repository"
)

type AuthUsecase interface {
	Login(r entity.LoginRequest) (entity.User, error)
	Register(r entity.RegisterRequest) (entity.User, error)
	CheckEmailExist(emailReq entity.RegisterRequest) (bool, error)
}

type authUsecase struct {
	authRepository repository.AuthRepository
}

func NewAuthUsecase(authRepository repository.AuthRepository) *authUsecase {
	return &authUsecase{
		authRepository: authRepository,
	}
}

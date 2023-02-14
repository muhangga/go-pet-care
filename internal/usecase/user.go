package usecase

import (
	"github.com/muhangga/internal/entity"
	"github.com/muhangga/internal/repository"
)

type UserUsecase interface {
	FetchAllUser() ([]entity.User, error)
	UpdateUser(entity entity.UserUpdateRequest, idUser int64) (entity.User, error)
	FindUserID(id int64) (entity.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) *userUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

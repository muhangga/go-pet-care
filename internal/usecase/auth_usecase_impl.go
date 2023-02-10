package usecase

import (
	"errors"
	"time"

	"github.com/muhangga/internal/entity"
	"github.com/muhangga/internal/utils"
)

func (a *authUsecase) Login(r entity.LoginRequest) (entity.User, error) {
	email := r.Email
	password := r.Password

	user, err := a.authRepository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	// if user.ID == sql.NullInt64{ false} {
	// 	return user, errors.New("email not found")
	// }

	matches := utils.ComparePassword(password, user.Password)
	if !matches {
		return user, errors.New("invalid password")
	}

	return user, nil
}

func (a *authUsecase) Register(r entity.RegisterRequest) (entity.User, error) {

	user := entity.User{}
	user.FullName = r.FullName
	user.Email = r.Email

	passwordHash, err := utils.GeneratePassword(r.Password)
	if err != nil {
		return user, err
	}
	user.Password = passwordHash
	user.Role = "user"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// valid, err := a.authRepository.FindByEmail(user.Email)
	// if err != nil {
	// 	return user, err
	// }

	// if valid.ID != 0 {
	// 	return user, errors.New("email already exists")
	// }

	user, err = a.authRepository.Save(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

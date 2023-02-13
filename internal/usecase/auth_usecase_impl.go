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

	if email == "" {
		return entity.User{}, errors.New("email is required")
	}

	if password == "" {
		return entity.User{}, errors.New("password is required")
	}

	user, err := a.authRepository.FindByEmail(email)
	if err != nil {
		return user, errors.New("email not found")
	}

	err = utils.ComparePassword(user.Password, password)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (a *authUsecase) Register(r entity.RegisterRequest) (entity.User, error) {

	user := entity.User{}

	if r.FullName == "" {
		return user, errors.New("full name is required")
	}

	if r.Email == "" {
		return user, errors.New("email is required")
	}

	if r.Password == "" {
		return user, errors.New("password is required")
	}

	if len(r.Password) < 6 {
		return user, errors.New("password must be 6 characters")
	}

	user.FullName = r.FullName
	user.Email = r.Email

	passwordHash, err := utils.GeneratePassword(r.Password)
	if err != nil {
		return user, err
	}
	user.Password = passwordHash
	user.Avatar = "avatar.jpg"
	user.Role = "user"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	user, err = a.authRepository.Save(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (a *authUsecase) CheckEmailExist(emailReq entity.RegisterRequest) (bool, error) {
	email := emailReq.Email

	user, err := a.authRepository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return false, nil
	}

	return true, nil
}

package usecase

import (
	"errors"
	"time"

	"github.com/muhangga/internal/entity"
)

func (u *userUsecase) FetchAllUser() ([]entity.User, error) {
	users, err := u.userRepository.GetAllUser()
	if err != nil {
		return users, nil
	}

	return users, nil
}

func (u *userUsecase) UpdateUser(e entity.UserUpdateRequest, idUser int64) (entity.User, error) {
	if e.FullName == "" {
		return entity.User{}, errors.New("full name required")
	}

	if *e.Gender == "" {
		return entity.User{}, errors.New("gender required")
	}

	if e.Email == "" {
		return entity.User{}, errors.New("email required")
	}

	user, err := u.userRepository.FindByID(idUser)
	if err != nil {
		return user, err
	}

	user.ID = idUser
	user.FullName = e.FullName
	user.Email = e.Email
	user.Gender = e.Gender
	user.Avatar = e.Avatar
	user.AboutMe = e.AboutMe
	user.UpdatedAt = time.Now()

	updateUser, err := u.userRepository.UpdateUser(user)
	if err != nil {
		return updateUser, err
	}

	return updateUser, nil
}

func (u *userUsecase) FindUserID(id int64) (entity.User, error) {
	if id == 0 {
		return entity.User{}, errors.New("id not found")
	}

	userID, err := u.userRepository.FindByID(id)
	if err != nil {
		return entity.User{}, err
	}

	return userID, nil
}

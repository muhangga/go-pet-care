package usecase

import (
	"errors"
	"time"

	"github.com/muhangga/internal/entity"
)

func (u *categoryUsecase) Save(c entity.CategoryRequest) (entity.Category, error) {
	category := entity.Category{}

	if c.Name == "" {
		return entity.Category{}, errors.New("name is required")
	}

	if c.Icon == "" {
		return entity.Category{}, errors.New("icon is required")
	}

	if c.Name == category.Name {
		return entity.Category{}, errors.New("name is already exist")
	}

	category.Name = c.Name
	category.Icon = c.Icon
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()
	return u.categoryRepository.Save(category)
}

func (u *categoryUsecase) FindAll() ([]entity.Category, error) {
	return u.categoryRepository.FindAll()
}

func (u *categoryUsecase) Update(c entity.CategoryRequest) (entity.Category, error) {
	category := entity.Category{}

	if c.Name == "" {
		return entity.Category{}, errors.New("name is required")
	}

	if c.Icon == "" {
		return entity.Category{}, errors.New("icon is required")
	}

	if c.Name == category.Name {
		return entity.Category{}, errors.New("name is already exist")
	}

	category.ID = c.ID
	category.Name = c.Name
	category.Icon = c.Icon
	category.UpdatedAt = time.Now()
	return u.categoryRepository.Update(category)
}

func (u *categoryUsecase) Delete(id int64) error {
	return u.categoryRepository.Delete(id)
}

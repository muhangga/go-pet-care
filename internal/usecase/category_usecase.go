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
	category, err := u.categoryRepository.FindAll()
	if err != nil {
		return nil, errors.New("error when find all category")
	}
	return category, nil
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
	if err := u.categoryRepository.Delete(id); err != nil {
		return errors.New("error when delete category")
	}
	return nil
}

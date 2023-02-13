package usecase

import (
	"github.com/muhangga/internal/entity"
	"github.com/muhangga/internal/repository"
)

type CategoryUsecase interface {
	Save(c entity.CategoryRequest) (entity.Category, error)
	FindAll() ([]entity.Category, error)
	Update(c entity.CategoryRequest) (entity.Category, error)
	Delete(id int64) error
}

type categoryUsecase struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryUsecase(categoryRepository repository.CategoryRepository) *categoryUsecase {
	return &categoryUsecase{
		categoryRepository: categoryRepository,
	}
}

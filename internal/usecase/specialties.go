package usecase

import (
	"github.com/muhangga/internal/entity"
	"github.com/muhangga/internal/repository"
)

type SpecialtiesUsecase interface {
	Save(s entity.SpecialtiesRequest) (entity.Specialties, error)
	FindAll() ([]entity.Specialties, error)
}

type specialtiesUsecase struct {
	specialtiesRepository repository.SpecialtiesRepository
}

func NewSpecialtiesUsecase(specialtiesRepository repository.SpecialtiesRepository) *specialtiesUsecase {
	return &specialtiesUsecase{
		specialtiesRepository: specialtiesRepository,
	}
}

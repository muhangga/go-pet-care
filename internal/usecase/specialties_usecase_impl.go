package usecase

import (
	"errors"
	"time"

	"github.com/muhangga/internal/entity"
)

func (u *specialtiesUsecase) Save(s entity.SpecialtiesRequest) (entity.Specialties, error) {

	specialties := entity.Specialties{}

	if s.Name == "" {
		return entity.Specialties{}, errors.New("name is required")
	}

	specialties.Name = s.Name
	specialties.Popular = true
	specialties.CreatedAt = time.Now()
	specialties.UpdatedAt = time.Now()

	return u.specialtiesRepository.Save(specialties)
}

func (u *specialtiesUsecase) FindAll() ([]entity.Specialties, error) {
	specialties, err := u.specialtiesRepository.FindAllSpecialties()
	if err != nil {
		return nil, errors.New("error when find all specialties")
	}
	return specialties, nil
}

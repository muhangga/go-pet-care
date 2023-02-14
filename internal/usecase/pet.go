package usecase

import (
	"github.com/muhangga/internal/entity"
	"github.com/muhangga/internal/repository"
)

type PetUsecase interface {
	SavePetAdditional(u entity.PetRequest) (entity.Pet, error)
}

type petUsecase struct {
	petRepository repository.PetRepository
}

func NewPetUsecase(petRepository repository.PetRepository) *petUsecase {
	return &petUsecase{petRepository: petRepository}
}

package usecase

import (
	"errors"
	"time"

	"github.com/muhangga/internal/entity"
)

func (u *petUsecase) SavePetAdditional(req entity.PetRequest) (entity.Pet, error) {
	if req.PetName == "" {
		return entity.Pet{}, errors.New("pet name required")
	}

	if req.Gender == "" {
		return entity.Pet{}, errors.New("gender required")
	}

	pet := entity.Pet{
		UserID:                 1,
		PetName:                req.PetName,
		Gender:                 req.Gender,
		DateOfBirth:            req.DateOfBirth,
		Neureted:               req.Neureted,
		Vaccinated:             req.Vaccinated,
		FriendlyWithDogs:       req.FriendlyWithDogs,
		FriendlyWithCats:       req.FriendlyWithCats,
		FriendlyWithKidsOver10: req.FriendlyWithKidsOver10,
		Microchipped:           req.Microchipped,
		Purebred:               req.Purebred,
		PetNurseryName:         req.PetNurseryName,
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
	}

	save, err := u.petRepository.Save(pet)
	if err != nil {
		return pet, err
	}

	return save, nil
}

func (u *petUsecase) GetAllPet() ([]entity.Pet, error) {
	pet, err := u.petRepository.GetAll()
	if err != nil {
		return pet, err
	}
	return pet, nil
}
func (u *petUsecase) FindByID(id int64) (entity.Pet, error) {
	if id == 0 {
		return entity.Pet{}, errors.New("ID not found")
	}

	petID, err := u.petRepository.FindPetByID(id)
	if err != nil {
		return petID, err
	}
	return petID, nil
}

package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhangga/internal/entity"
	"github.com/muhangga/internal/usecase"
	"github.com/muhangga/internal/utils"
)

type petDelivery struct {
	petUsecase usecase.PetUsecase
}

func NewPetDelivery(petUsecase usecase.PetUsecase) *petDelivery {
	return &petDelivery{
		petUsecase: petUsecase,
	}
}

func (d *petDelivery) CreatePet(c *gin.Context) {
	var petRequest entity.PetRequest

	if err := c.ShouldBindJSON(&petRequest); err != nil {
		formatter := utils.ErrorResponse("Something when wrong", http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, formatter)
		return
	}

	pet, err := d.petUsecase.SavePetAdditional(petRequest)
	if err != nil {
		formatter := utils.ErrorResponse("Failed to create pet detail", http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, formatter)
		return
	}

	formatter := utils.SuccessResponse("Success to create pet information", http.StatusOK, pet)
	c.JSON(http.StatusOK, formatter)
}

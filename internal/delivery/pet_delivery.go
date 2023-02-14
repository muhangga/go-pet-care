package delivery

import (
	"net/http"
	"strconv"

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

func (d *petDelivery) FetchAllPet(c *gin.Context) {
	pet, err := d.petUsecase.GetAllPet()
	if err != nil {
		formatter := utils.ErrorResponse("Failed to load all pet", http.StatusBadRequest, err.Error())

		c.JSON(http.StatusBadRequest, formatter)
		return
	}

	formatter := utils.SuccessResponse("Fetch All Pet", http.StatusOK, pet)
	c.JSON(http.StatusOK, formatter)
}

func (d *petDelivery) FindPetByID(c *gin.Context) {
	id := c.Param("id")

	convertID, _ := strconv.Atoi(id)
	pet, err := d.petUsecase.FindByID(int64(convertID))

	if err != nil {
		formatter := utils.ErrorResponse("failed to load pet detail", http.StatusBadRequest, err.Error())

		c.JSON(http.StatusBadRequest, formatter)
		return
	}

	formatter := utils.SuccessResponse("Success to load pet detail", http.StatusOK, pet)
	c.JSON(http.StatusOK, formatter)
}

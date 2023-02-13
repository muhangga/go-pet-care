package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhangga/internal/entity"
	"github.com/muhangga/internal/usecase"
	"github.com/muhangga/internal/utils"
)

type specialtiesDelivery struct {
	specialtiesUsecase usecase.SpecialtiesUsecase
}

func NewSpecialtiesDelivery(specialtiesUsecase usecase.SpecialtiesUsecase) *specialtiesDelivery {
	return &specialtiesDelivery{
		specialtiesUsecase: specialtiesUsecase,
	}
}

func (d *specialtiesDelivery) CreateSpecialties(c *gin.Context) {

	var specialtiesRequest entity.SpecialtiesRequest

	if err := c.ShouldBindJSON(&specialtiesRequest); err != nil {
		formatter := utils.ErrorResponse("Create specialties failed", http.StatusBadRequest, err.Error())

		c.JSON(http.StatusBadRequest, formatter)
		return
	}

	specialties, err := d.specialtiesUsecase.Save(specialtiesRequest)
	if err != nil {
		formatter := utils.ErrorResponse("Create specialties failed", http.StatusBadRequest, err.Error())

		c.JSON(http.StatusBadRequest, formatter)
		return
	}

	formatter := utils.SuccessResponse("Create specialties success", http.StatusOK, specialties)
	c.JSON(http.StatusOK, formatter)
}

func (d *specialtiesDelivery) FetchAllSpecialties(c *gin.Context) {
	specialties, err := d.specialtiesUsecase.FindAll()
	if err != nil {
		formatter := utils.ErrorResponse("Failed to load specialties", http.StatusBadRequest, err.Error())

		c.JSON(http.StatusBadRequest, formatter)
		return
	}

	formatter := utils.SuccessResponse("Fetch All Specialties", http.StatusOK, specialties)
	c.JSON(http.StatusOK, formatter)
}

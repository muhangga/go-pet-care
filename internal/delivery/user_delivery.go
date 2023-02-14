package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhangga/internal/entity"
	"github.com/muhangga/internal/usecase"
	"github.com/muhangga/internal/utils"
)

type userDelivery struct {
	userUsecase usecase.UserUsecase
}

func NewUserDelivery(userUsecase usecase.UserUsecase) *userDelivery {
	return &userDelivery{
		userUsecase: userUsecase,
	}
}

func (d *userDelivery) UpdateProfileUser(c *gin.Context) {
	var userRequest entity.UserUpdateRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		formatter := utils.ErrorResponse("Failed to update user", http.StatusBadRequest, err.Error())

		c.JSON(http.StatusBadRequest, formatter)
		return
	}

	currentUser := c.MustGet("currentUser").(entity.User)

	updateUser, err := d.userUsecase.UpdateUser(userRequest, int64(currentUser.ID))
	if err != nil {
		formatter := utils.ErrorResponse("Failed to update userr", http.StatusBadRequest, err.Error())

		c.JSON(http.StatusBadRequest, formatter)
		return
	}

	formatter := utils.SuccessResponse("Update user success", http.StatusOK, updateUser)
	c.JSON(http.StatusOK, formatter)
}

package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhangga/internal/entity"
	"github.com/muhangga/internal/usecase"
	"github.com/muhangga/internal/utils"
	"github.com/muhangga/pkg/middleware"
)

type authDelivery struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthDelivery(authUsecase usecase.AuthUsecase) authDelivery {
	return authDelivery{authUsecase: authUsecase}
}

func (d *authDelivery) Login(c *gin.Context) {
	var loginRequest entity.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		formatter := utils.ErrorResponse("Login failed", http.StatusBadRequest, err.Error())

		c.JSON(http.StatusBadRequest, formatter)
		return
	}

	user, err := d.authUsecase.Login(loginRequest)
	if err != nil {
		formatter := utils.ErrorResponse("Login failed", http.StatusBadRequest, err.Error())

		c.JSON(http.StatusBadRequest, formatter)
		return
	}

	// generate token JWT
	token, err := middleware.GenerateToken(user.ID, user.Email)
	if err != nil {
		formatter := utils.ErrorResponse("Login failed", http.StatusBadRequest, err.Error())

		c.JSON(http.StatusBadRequest, formatter)
		return
	}

	loginSuccessFormatter := utils.FormatLoginSuccess(user, token)

	formatter := utils.SuccessResponse("Login success", http.StatusOK, loginSuccessFormatter)
	c.JSON(http.StatusOK, formatter)
}

func (d *authDelivery) Register(c *gin.Context) {
	var registerRequest entity.RegisterRequest

	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		formatter := utils.ErrorResponse("Register failed", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, formatter)
		return
	}

	isExist, _ := d.authUsecase.CheckEmailExist(registerRequest)
	if isExist {
		formatter := utils.ErrorResponse("Register failed", http.StatusBadRequest, "email already exist")
		c.JSON(http.StatusBadRequest, formatter)

		return
	}

	user, err := d.authUsecase.Register(registerRequest)
	if err != nil {
		formatter := utils.ErrorResponse("Register failed", http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, formatter)

		return
	}

	formatter := utils.UserFormatter(user)
	response := utils.SuccessResponse("Account has been registered", http.StatusOK, formatter)
	c.JSON(http.StatusOK, response)
}

package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhangga/internal/entity"
	"github.com/muhangga/internal/usecase"
	"github.com/muhangga/internal/utils"
)

type authDelivery struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthDelivery(authUsecase usecase.AuthUsecase) authDelivery {
	return authDelivery{authUsecase: authUsecase}
}

func (a *authDelivery) Login(c *gin.Context) {
	var loginRequest entity.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := a.authUsecase.Login(loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})

	// token, err := utils.GenerateToken(user)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{"token": token})
}

func (a *authDelivery) Register(c *gin.Context) {
	var registerRequest entity.RegisterRequest

	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := a.authUsecase.Register(registerRequest)
	if err != nil {
		formatter := utils.ErrorResponse("Register failed", http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, formatter)
		return
	}

	formatter := utils.UserFormatter(user)

	response := utils.SuccessResponse("Account has been registered", http.StatusOK, formatter)
	c.JSON(http.StatusOK, response)
}

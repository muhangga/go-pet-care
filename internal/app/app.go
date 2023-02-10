package app

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muhangga/config"
	"github.com/muhangga/internal/delivery"
	"github.com/muhangga/internal/repository"
	"github.com/muhangga/internal/usecase"
)

type server struct {
	httpServer *gin.Engine
	config     config.Config
}

type Server interface {
	Run()
}

func InitServer(config config.Config) Server {
	return &server{
		httpServer: gin.Default(),
		config:     config,
	}
}

func (s *server) Run() {

	authRepository := repository.NewRepository(s.config.Database())
	authUsecase := usecase.NewUserUsecase(authRepository)
	authDelivery := delivery.NewAuthDelivery(authUsecase)

	api := s.httpServer.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/login", authDelivery.Login)
		auth.POST("/register", authDelivery.Register)
	}

	if err := s.httpServer.Run(":" + strconv.Itoa(s.config.ServicePort())); err != nil {
		log.Panic(err)
	}
}

package app

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muhangga/config"
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

	s.httpServer.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})

	})

	if err := s.httpServer.Run(":" + strconv.Itoa(s.config.ServicePort())); err != nil {
		log.Panic(err)
	}
}

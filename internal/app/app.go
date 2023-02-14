package app

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/muhangga/config"
	"github.com/muhangga/internal/delivery"
	"github.com/muhangga/internal/repository"
	"github.com/muhangga/internal/usecase"
	"github.com/muhangga/pkg/middleware"
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

func (s *server) DB() *sql.DB {
	return s.config.Database()
}

func (s *server) Run() {

	middleware := middleware.InitMiddleware()
	s.httpServer.Use(middleware.CORS())

	cookieStore := cookie.NewStore([]byte(os.Getenv("JWT_TOKEN")))
	s.httpServer.Use(sessions.Sessions("go-pet", cookieStore))

	authRepository := repository.NewAuthRepository(s.DB())
	authUsecase := usecase.NewAuthUsecase(authRepository)
	authDelivery := delivery.NewAuthDelivery(authUsecase)

	categoryRepository := repository.NewCategoryRepository(s.DB())
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepository)
	categoryDelivery := delivery.NewCategoryDelivery(categoryUsecase)

	specialtiesRepository := repository.NewSpecialtiesRepository(s.DB())
	specialtiesUsecase := usecase.NewSpecialtiesUsecase(specialtiesRepository)
	specialtiesDelivery := delivery.NewSpecialtiesDelivery(specialtiesUsecase)

	petRepository := repository.NewPetRepository(s.DB())
	petUsecase := usecase.NewPetUsecase(petRepository)
	petDelivery := delivery.NewPetDelivery(petUsecase)

	userRepository := repository.NewUserRepository(s.DB())
	userUsecase := usecase.NewUserUsecase(userRepository)
	userDelivery := delivery.NewUserDelivery(userUsecase)

	api := s.httpServer.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/login", authDelivery.Login)
		auth.POST("/register", authDelivery.Register)
	}

	category := api.Group("/category").Use(middleware.JWTMiddleware(userUsecase))
	{
		category.POST("/create", categoryDelivery.CreateCategory)
		category.GET("/all", categoryDelivery.FetchAllCategory)
		category.PUT("/update", categoryDelivery.UpdateCategory)
		category.DELETE("/delete/:id", categoryDelivery.DeleteCategory)
	}

	specialties := api.Group("/specialties").Use(middleware.JWTMiddleware(userUsecase))
	{
		specialties.POST("/create", specialtiesDelivery.CreateSpecialties)
		specialties.GET("/all", specialtiesDelivery.FetchAllSpecialties)
	}

	pet := api.Group("/pet").Use(middleware.JWTMiddleware(userUsecase))
	{
		pet.POST("/create", petDelivery.CreatePet)
		pet.GET("/fetch", petDelivery.FetchAllPet)
		pet.GET("/fetch/:id", petDelivery.FindPetByID)
	}

	user := api.Group("/user").Use(middleware.JWTMiddleware(userUsecase))
	{
		user.PUT("/update/:id", userDelivery.UpdateProfileUser)
	}

	if err := s.httpServer.Run(":" + strconv.Itoa(s.config.ServicePort())); err != nil {
		log.Panic(err)
	}
}

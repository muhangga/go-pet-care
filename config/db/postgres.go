package db

import (
	"fmt"
	"os"

	"github.com/muhangga/internal/entity"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := os.Getenv("POSTGRES_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error().Msgf("cant connect to database %s", err)
	}

	db.AutoMigrate(&entity.User{}, &entity.Specialties{}, &entity.Doctor{}, &entity.Pet{}, &entity.Category{}, &entity.Booking{}, &entity.Additional{})
	fmt.Println("Database connected")

	return db
}

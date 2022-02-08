package database

import (
	"go-forum-thingy/models"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	dsn := "host=localhost user=admin password=PASSWORD dbname=DBNAME port=5432 sslmode=disable"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal().Msg("Failed to connect to database")
	}

	DB = connection
	connection.AutoMigrate(&models.User{})
}


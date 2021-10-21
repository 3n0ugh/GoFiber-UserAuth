package database

import (
	"fmt"
	"log"
	"os"

	"github.com/3n0ugh/GoFiber-RestAPI-UserAuth/server/config"
	"github.com/3n0ugh/GoFiber-RestAPI-UserAuth/server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	// get values from dotenv file
	config, err := config.GetConfig()
	if err != nil {
		log.Fatal("Can't get config", err.Error())
		os.Exit(1)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s", config.DatabaseHost, config.DatabaseUsername, config.DatabasePassword, config.DatabaseName, config.DatabasePort, config.DatabaseTimeZone)
	// database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database, \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	// add logger
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	// enable auto migration
	db.AutoMigrate(&models.User{})

	DB = Dbinstance{
		Db: db,
	}
}

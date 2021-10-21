package database

import (
	"fmt"
	"log"
	"os"

	"github.com/3n0ugh/GoFiber-RestAPI-UserAuth/server/models"
	"github.com/golobby/dotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

type Config struct {
	Name     string `env:"DB_NAME"`
	Port     int16  `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	Pass     string `env:"DB_PASS"`
	Host     string `env:"DB_HOST"`
	TimeZone string `env:"DB_TIMEZONE"`
}

func ConnectDb() {
	config := &Config{}
	file, err := os.Open(".env")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = dotenv.NewDecoder(file).Decode(config)
	if err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s", config.Host, config.User, config.Pass, config.Name, config.Port, config.TimeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database, \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	db.AutoMigrate(&models.User{}) // Here is the model &models.User{}

	DB = Dbinstance{
		Db: db,
	}
}

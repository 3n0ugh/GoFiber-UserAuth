package config

import (
	"os"
	"time"

	"github.com/golobby/dotenv"
)

type Config struct {
	ServerUrl     string        `env:"SERVER_URL"`
	ServerTimeout time.Duration `env:"SERVER_TIMEOUT"`

	DatabaseName     string `env:"DATABASE_NAME"`
	DatabasePort     int16  `env:"DB_PORT"`
	DatabaseUsername string `env:"DB_USER"`
	DatabasePassword string `env:"DB_PASS"`
	DatabaseHost     string `env:"DB_HOST"`
	DatabaseTimeZone string `env:"DB_TIMEZONE"`

	JwtSecretKey  string        `env:"JWT_SECRET_KEY"`
	JwtExpireTime time.Duration `env:"JWT_EXPIRE_TIME"`
}

func GetConfig() (*Config, error) {
	// initialize the config struct
	config := &Config{}

	// open dotenv file
	file, err := os.Open(".env")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// adding data to config struct
	err = dotenv.NewDecoder(file).Decode(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

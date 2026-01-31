package utils

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)


type Config struct {
	Name        string `env:"APP_NAME" envDefault:"pace-backend"`
    HttpPort    int    `env:"HTTP_PORT" envDefault:"8080"`
    LogLevel    string `env:"LOG_LEVEL" envDefault:"info"`
    Environment string `env:"ENVIRONMENT" envDefault:"local"`

	PostgresDBUri string `env:"POSTGRES_DB_URI"`
	
	FirebaseProjectID	string `env:"FIREBASE_PROJECT_ID"`
	FirebaseClientEmail	string `env:"FIREBASE_CLIENT_EMAIL"`
	FirebasePrivateKey	string `env:"FIREBASE_PRIVATE_KEY"`
}

var appConfig *Config

func GetConfig() *Config {
	if  appConfig != nil {
		return appConfig
	}else {
		err := godotenv.Load(".env")
		if err != nil {
			fmt.Println("Unable to load .env file")
		}
		appConfig = &Config{}
		if err = env.Parse(appConfig); err != nil {
			panic(err)
		}
		return appConfig
	}
}

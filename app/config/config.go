package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	DBName          string
	ValidateUserAdd string
	AppUrl          string
}

func LoadFromEnv() DBConfig {
	var config DBConfig
	if runningInGitHubWorkflow() {
		temp_port, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
		config = DBConfig{
			Host:            os.Getenv("POSTGRES_HOST"),
			Port:            temp_port,
			User:            os.Getenv("POSTGRES_USER"),
			Password:        os.Getenv("POSTGRES_PASSWORD"),
			DBName:          os.Getenv("POSTGRES_DB"),
			ValidateUserAdd: os.Getenv("VALIDATE_USER_ADDRESS"),
			AppUrl:          os.Getenv("APP_URL"),
		}
	} else {
		err := godotenv.Load("./.env")
		if err != nil {
			log.Fatal("Failed to load the config")
		}
		temp_port, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
		config = DBConfig{
			Host:            os.Getenv("POSTGRES_HOST"),
			Port:            temp_port,
			User:            os.Getenv("POSTGRES_USER"),
			Password:        os.Getenv("POSTGRES_PASSWORD"),
			DBName:          os.Getenv("POSTGRES_DB"),
			ValidateUserAdd: os.Getenv("VALIDATE_USER_ADDRESS"),
			AppUrl:          os.Getenv("APP_URL"),
		}
	}
	return config
}

func runningInGitHubWorkflow() bool {
	return os.Getenv("CI") == "true"
}

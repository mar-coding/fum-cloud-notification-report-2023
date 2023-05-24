package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func Connect() (*sql.DB, error) {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Failed to load the config")
	}
	temp_port, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	config := DBConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     temp_port,
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", config.Host, config.User, config.Password, config.DBName, config.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected Successfully to the Database")
	return sqlDB, nil
}

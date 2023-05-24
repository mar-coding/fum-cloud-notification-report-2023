package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/mar-coding/fum-cloud-notification-report-2023/app/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*sql.DB, error) {
	config := config.LoadFromEnv()
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

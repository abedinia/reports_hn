package db

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"report_hn/internal/config"
	"report_hn/internal/logger"
)

var dbInstance *gorm.DB

func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Log.Fatal("Failed to connect to database: ", err)
		return nil, err
	}

	fmt.Println("Successfully connected to PostgreSQL database!")
	dbInstance = db
	return db, nil
}

func GetDB() *gorm.DB {
	dbConfig := config.AppConfig.Database
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Dbname)

	if dbInstance == nil {
		InitDB(connStr)
	}
	return dbInstance
}

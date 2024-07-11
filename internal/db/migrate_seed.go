package db

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"report_hn/internal/config"
	"report_hn/internal/logger"
)

func RunMigrations() {
	dbConfig := config.AppConfig.Database
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Dbname)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		logger.Log.Fatal(err)
	}

	if err := db.AutoMigrate(User{}, Report{}); err != nil {
		logger.Log.Fatal(err)
	}

	fmt.Println("Database migrations applied successfully!")
}

func SeedUser() {
	dbConfig := config.AppConfig.Database
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Dbname)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		logger.Log.Fatal(err)
	}

	username := "blog_bot"
	rawPassword := "123456"

	var existingUser User
	if err := db.Where("username = ?", username).First(&existingUser).Error; err == nil {
		logger.Log.Println("User already exists, skipping seeding")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		logger.Log.Fatalf("failed to hash password: %v", err)
	}

	user := User{
		Username: username,
		Password: string(hashedPassword),
	}

	if err := db.Create(&user).Error; err != nil {
		logger.Log.Fatalf("failed to create user: %v", err)
	}

	logger.Log.Println("User seeded successfully")
}

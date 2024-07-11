package db

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Report struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	DatasetID uint   `gorm:"not null"`
	Name      string `gorm:"not null"`
	Type      uint   `gorm:"not null"`
	Reason    string `gorm:"not null"`
	Handled   bool   `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

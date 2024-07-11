package server

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"report_hn/internal/token"
)

type User struct {
	ID       uint   `gorm:"primary_key"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func AuthLogin(psql *gorm.DB, c *gin.Context) {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var user User
	if err := psql.Where("username = ?", request.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	tokenString, err := token.MakeToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

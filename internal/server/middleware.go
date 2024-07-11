package server

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"report_hn/internal/token"
	"strings"
)

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		c.Abort()
		return
	}

	tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))
	userID, err := token.DecodeToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	c.Set("user_id", userID)
	c.Next()
}

func RateLimiter(limit rate.Limit, burst int) gin.HandlerFunc {
	limiter := rate.NewLimiter(limit, burst)
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.AbortWithStatusJSON(429, gin.H{
				"error": "Too many requests",
			})
			return
		}
		c.Next()
	}
}

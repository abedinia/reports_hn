package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"report_hn/internal/db"
	"report_hn/internal/logger"
	"strconv"
	"time"
)

func CreateReport(psql *gorm.DB, c *gin.Context) {
	var request struct {
		DatasetID uint   `json:"dataset_id"`
		Name      string `json:"name"`
		Type      uint   `json:"type"`
		Reason    string `json:"reason"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	report := db.Report{
		UserID:    userID.(uint),
		DatasetID: request.DatasetID,
		Name:      request.Name,
		Type:      request.Type,
		Reason:    request.Reason,
		Handled:   false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result := psql.Create(&report)
	if result.Error != nil {
		logger.Log.Error(result.Error)
		c.JSON(http.StatusNotFound, gin.H{"error": "some problem on db"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Report created successfully"})
}

func GetReports(psql *gorm.DB, c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "10")

	pageNum, err := strconv.Atoi(page)
	if err != nil || pageNum < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	pageSizeNum, err := strconv.Atoi(pageSize)
	if err != nil || pageSizeNum < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
		return
	}

	offset := (pageNum - 1) * pageSizeNum

	var reports []db.Report

	result := psql.Where("user_id = ?", userID).Offset(offset).Limit(pageSizeNum).Find(&reports)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "some problem on db"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reports": reports})
}

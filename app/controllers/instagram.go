package controllers

import (
	"errors"
	"github.com/7leven7/claimd/app/db"
	"github.com/7leven7/claimd/app/models"
	_ "github.com/7leven7/claimd/docs"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetInstagramData @BasePath /api/v1
// @Summary GetInstagramData
// @Schemes
// @Description  Retrieve all the Instagram data from the database
// @Success 200 {object} models.Instagram
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /instagram [get]
func GetInstagramData() gin.HandlerFunc {
	return func(c *gin.Context) {
		var instagram []models.Instagram
		username := "instagram"
		result := db.DB.Where("username = ?", username).Find(&instagram)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				c.JSON(404, gin.H{"error": "No record found!"})
				return
			}
			c.JSON(500, gin.H{"error": result.Error})
			return
		}
		c.JSON(200, instagram)
	}
}

package controllers

import (
	"github.com/7leven7/claimd/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ScrapePlatform @BasePath /api/v1
// @Summary ScrapePlatform
// @Schemes
// @Description Scrape all platforms
// @Produce json
// @Success 200 {object} string
// @Router /scrape [get]
func ScrapePlatform() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := services.Scrape()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Scraping successful"})
	}
}

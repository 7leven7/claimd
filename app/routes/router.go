package routes

import (
	"github.com/7leven7/claimd/app/controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InstagramRoutes sets up the routes
func InstagramRoutes(router *gin.Engine) {
	router.GET("/instagram", controllers.GetInstagramData())
}

// TiktokRoutes sets up the routes
func TiktokRoutes(router *gin.Engine) {
	router.GET("/tiktok", controllers.GetTiktokData())
}

// ScraperRoutes sets up the routes
func ScraperRoutes(router *gin.Engine) {
	router.GET("/scrape", controllers.ScrapePlatform())
}

// SwaggerRoutes sets up the routes
func SwaggerRoutes(router *gin.Engine) {
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

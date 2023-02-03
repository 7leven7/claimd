package main

import (
	"github.com/7leven7/claimd/app/configs"
	"github.com/7leven7/claimd/app/routes"
	"github.com/7leven7/claimd/app/seeders"
	"github.com/gin-gonic/gin"
)

// @title Claimd API
// @description This is a sample server for a claimd API.
// @version 1.0.0

func main() {

	err := configs.PGConnection()
	if err != nil {
		return
	}
	configs.Migrate()
	seeders.PlatformsSeeder()
	router := gin.Default()
	routes.InstagramRoutes(router)
	routes.TiktokRoutes(router)
	routes.SwaggerRoutes(router)
	routes.ScraperRoutes(router)

	router.Run()
}

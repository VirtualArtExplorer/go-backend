package main

import (
	"log"

	"github.com/VirtualArtExplorer/go-backend/controllers"
	"github.com/VirtualArtExplorer/go-backend/database"
	"github.com/VirtualArtExplorer/go-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func initConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	viper.SetConfigType("env")
	viper.AutomaticEnv()
}

func main() {
	initConfig()
	database.InitDB()
	database.RunMigrations()

	r := gin.Default()

	r.POST("/api/v1/managers", controllers.CreateManager)
	r.POST("/api/v1/login", controllers.Login)
	r.GET("/api/v1/museums/", controllers.GetAllMuseum)
	r.GET("api/v1/museums/city/:city", controllers.GetMuseumsByCity)
	r.GET("api/v1/museums/state/:state", controllers.GetMuseumsByState)
	r.GET("api/v1/museums/name/:name", controllers.GetMuseumsByName)

	auth := r.Group("/api/v1")
	auth.Use(utils.ValidateTokenMiddleware)

	auth.PUT("/managers/:id", controllers.UpdateManager)
	auth.PUT("/managers/:id/disable", controllers.DisableManager)
	auth.POST("/museums/", controllers.CreateMuseum)
	auth.PUT("/museums/:id", controllers.UpdateMuseum)
	auth.PUT("/museums/:id/disable", controllers.DisableMuseum)

	r.Run(":8080")
}

package main

import (
	"log"

	"gin-user-api/config"
	"gin-user-api/internal/handlers"
	"gin-user-api/internal/models"
	"gin-user-api/internal/repositories"
	"gin-user-api/internal/routes"
	"gin-user-api/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db := config.ConnectDB()

	db.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.Post{},
	)

	userHandler := handlers.NewUserHandler(db)
	profileHandler := handlers.NewProfileHandler(db)


	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	routes.RegisterRoutes(
		r,
		userHandler,
		profileHandler,
	)

	log.Println("Server running at :8081")
	r.Run(":8081")
}

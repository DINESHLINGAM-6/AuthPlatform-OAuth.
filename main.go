package main

import (
	"auth-system/database"
	"auth-system/handlers"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the database
	database.InitDB()

	// Set up the router
	r := gin.Default()

	// Routes
	r.POST("/signup", handlers.Signup)
	r.POST("/login", handlers.Login)
	r.GET("/protected", handlers.ProtectedRoute)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Server is running"})
	})

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	r.Run(":" + port)
}

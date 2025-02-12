package main

import (
	"auth-system/database"
	"auth-system/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	database.InitDB()

	// Set up the router
	r := gin.Default()

	// Routes
	r.POST("/signup", handlers.Signup)
	r.POST("/login", handlers.Login)

	// Protected route (example)
	r.GET("/protected", handlers.ProtectedRoute)

	// Start the server
	r.Run(":8080")
}
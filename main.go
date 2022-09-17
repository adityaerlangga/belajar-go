package main

import (
	"github.com/adityaerlangga/golang-auth/app/config"
	controllers "github.com/adityaerlangga/golang-auth/controllers/usercontroller"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connecting to database
	config.ConnectDatabase()

	r := gin.Default()
	// Routes
	r.GET("/users", controllers.Index)                   // All users
	r.GET("/user/:id", controllers.Show)                 // Get User By ID
	r.POST("/register", controllers.Register)            // Sign up / Register
	r.POST("/login", controllers.Login)                  // Sign in
	r.POST("/reset-password", controllers.ResetPassword) // Reset Password

	// Running server
	r.Run()
}

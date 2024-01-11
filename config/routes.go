package config

import (
	"log"
	"udala/sso/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	log.Println("Setting up controllers")
	router := gin.Default()

	// added CORS Middleware
	router.Use(CORSMiddleware())

	// controllers
	setupUserController(router, db)
	setupAuthController(router, db)

	return router
}

func setupUserController(router *gin.Engine, db *gorm.DB) {
	// creating router
	userController := controller.NewUserController(db)

	// adding routes
	router.POST("/users", userController.CreateUser)
	router.GET("/users", userController.GetUsers)
	router.GET("/users/:id", userController.GetUserByID)
	router.DELETE("/users/:id", userController.DeleteUser)
	router.GET("/users/check-availability", userController.CheckAvailability)
}

func setupAuthController(router *gin.Engine, db *gorm.DB) {
	authController := controller.NewAuthController(db)

	router.POST("/auth", authController.AuthLogin)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Allow all origin
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		// Allow specific HTTP methods
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")

		// Allow specific HTTP headers
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// Allow credentials (cookies, authentication)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle OPTIONS request (pre-flight request)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204) // No content for pre-flight requests
			return
		}

		c.Next()
	}
}

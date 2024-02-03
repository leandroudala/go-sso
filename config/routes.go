package config

import (
	"log"
	"udala/sso/controller"
	"udala/sso/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	log.Println("Setting up controllers")
	router := gin.Default()

	// added CORS Middleware
	router.Use(CORSMiddleware())

	// setup api version
	setupApiVersion1(router, db)

	return router
}

func setupApiVersion1(router *gin.Engine, db *gorm.DB) {
	v1BasePath := "/api/v1"
	v1 := router.Group(v1BasePath)

	docs.SwaggerInfo.BasePath = "/api/v1"

	// swagger
	setupSwagger(router)

	// controllers
	setupUserController(router, db, v1)
	setupAuthController(router, db, v1)
}

func setupSwagger(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func setupUserController(router *gin.Engine, db *gorm.DB, version *gin.RouterGroup) {
	// creating router
	userController := controller.NewUserController(db)

	api := version.Group("/users")

	// adding routes
	api.POST("/", userController.CreateUser)
	api.GET("/", userController.GetUsers)
	api.GET("/:id", userController.GetUserByID)
	api.DELETE("/:id", userController.DeleteUser)
	api.GET("/check-availability", userController.CheckAvailability)
}

func setupAuthController(router *gin.Engine, db *gorm.DB, version *gin.RouterGroup) {
	authController := controller.NewAuthController(db)

	api := version.Group("/auth")
	api.POST("/", authController.AuthLogin)
	api.POST("/forget-password", authController.ForgetPassword)
	api.POST("/confirm-email", authController.ConfirmEmail)
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

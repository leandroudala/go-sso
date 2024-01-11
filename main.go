package main

import (
	"log"
	"udala/sso/config"
	"udala/sso/database"
)

// @title Udala SSO
// @version 1.0
// @description Base for SSO application written in Go
// @BasePath /api/v1
func main() {
	// reading environment variables
	config.ReadEnv()

	// generate database structure
	db := database.Setup()

	// setup route
	router := config.SetupRouter(db)

	log.Println("Starting application")
	router.Run(config.GetAddr())
}

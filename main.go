package main

import (
	"log"
	"udala/sso/config"
	"udala/sso/database"
)

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

package database

import (
	"log"
	"udala/sso/config"
	"udala/sso/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetMySQLConnection() *gorm.DB {
	dsn := config.GetMySQLDatabaseDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func autoMigrate(db *gorm.DB) {
	log.Println("Starting migration")

	err := db.AutoMigrate(
		model.User{},
	)
	if err != nil {
		panic(err)
	}
}

func Setup() *gorm.DB {
	db := GetMySQLConnection()
	log.Println("Checking database")
	if config.IsAutoMigrateEnabled() {
		autoMigrate(db)
	}

	return db
}

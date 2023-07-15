package db

import (
	"aurora/services/db/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var Client *gorm.DB

func Init() {
	host, hostOk := os.LookupEnv("DB_HOST")
	user, userOk := os.LookupEnv("DB_USER")
	password, passwordOk := os.LookupEnv("DB_PASSWORD")
	dbname, dbnameOk := os.LookupEnv("DB_NAME")
	port, portOk := os.LookupEnv("DB_PORT")
	timezone, timezoneOk := os.LookupEnv("DB_TIMEZONE")

	if !hostOk || !userOk || !passwordOk || !dbnameOk || !portOk || !timezoneOk {
		panic("Missing database environment variables!")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", host, user, password, dbname, port, timezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	Client = db
}

func AutoMigrate() error {
	return Client.AutoMigrate(
		&models.Auth{},
		&models.User{},
		&models.Brand{},
		&models.Product{},
		&models.ProductImage{},
		&models.AdPreference{},
		&models.Address{},
	)
}

package db

import (
	"aurora/db/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

var Client *gorm.DB

func createDsnFromEnvVars() string {
	env := map[string]string{
		"DB_HOST":     "",
		"DB_USER":     "",
		"DB_PASSWORD": "",
		"DB_NAME":     "",
		"DB_PORT":     "",
		"DB_TIMEZONE": "",
	}

	for key := range env {
		value, ok := os.LookupEnv(key)

		if !ok {
			panic(fmt.Sprintf("Missing database environment variable %s", key))
		}

		env[key] = value
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		env["DB_HOST"],
		env["DB_USER"],
		env["DB_PASSWORD"],
		env["DB_NAME"],
		env["DB_PORT"],
		env["DB_TIMEZONE"],
	)

	return dsn
}

func getLogLevelFromEnv() logger.LogLevel {
	logLevel := logger.Silent
	debug := os.Getenv("DEBUG")

	if debug == "true" {
		logLevel = logger.Info
	}

	return logLevel
}

func Init() {
	dsn := createDsnFromEnvVars()
	logLevel := getLogLevelFromEnv()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})

	if err != nil {
		panic(err)
	}

	Client = db
}

func AutoMigrate() error {
	return Client.AutoMigrate(
		&models.Address{},
		&models.Admin{},
		&models.AdPreference{},
		&models.Auth{},
		&models.Brand{},
		&models.BrandReview{},
		&models.Cart{},
		&models.CartItem{},
		&models.Category{},
		&models.Favorite{},
		&models.Product{},
		&models.ProductVariant{},
		&models.ProductImage{},
		&models.ProductReview{},
		&models.ProductStyle{},
		&models.ProductSize{},
		&models.User{},
	)
}

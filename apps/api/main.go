package main

import (
	"aurora/db"
	"aurora/handlers"
	"aurora/middlewares"
	"aurora/services/tasks"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	if err := router.SetTrustedProxies(nil); err != nil {
		panic(err.Error())
	}

	router.Use(middlewares.Cors())

	// Call only once to initialize the database connection
	// Panics if the connection fails
	db.Init()

	// Call only once to make sure the database schema is up-to-date
	// All database entities are registered under this function
	if err := db.AutoMigrate(); err != nil {
		panic(err.Error())
	}

	// Call only once to initialize the tasks service
	// Need to run in a different goroutine
	// Serves as a background worker
	go tasks.Init()
	defer tasks.Close()

	// Creates a router group for the /api/v1 path
	// Attaches global middlewares.
	// Defines all routes.
	handlers.Bootstrap(router)

	if err := router.Run(":5000"); err != nil {
		panic(err.Error())
	}
}

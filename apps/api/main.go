package main

import (
	"aurora/db"
	"aurora/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	err := router.SetTrustedProxies(nil)
	router.Use(cors.Default())

	if err != nil {
		panic(err.Error())
	}

	// Call only once to initialize the database connection
	// Panics if the connection fails
	db.Init()

	err = db.AutoMigrate()

	if err != nil {
		panic(err.Error())
	}

	handlers.Bootstrap(router)

	err = router.Run(":5000")

	if err != nil {
		panic(err.Error())
	}
}

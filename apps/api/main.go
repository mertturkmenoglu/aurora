package main

import (
	"aurora/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.Bootstrap(router)

	err := router.Run(":3000")

	if err != nil {
		fmt.Println("Cannot run application")
	}
}

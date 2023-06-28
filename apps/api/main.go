package main

import (
	"aurora/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	err := router.SetTrustedProxies(nil)

	if err != nil {
		panic(err.Error())
	}

	routes.Bootstrap(router)

	err = router.Run(":3000")

	if err != nil {
		panic(err.Error())
	}
}

package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"submission2/routes"
)

func main() {
	router := gin.Default()
	routes.SetupRoutes(router)
	log.Fatal(router.Run(":8080"))
}
package main

import (
	"log"

	routes "github.com/GorginZ/community-library/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/books", routes.HandleBooks)

	log.Fatal(router.Run("0.0.0.0:8080"))
}

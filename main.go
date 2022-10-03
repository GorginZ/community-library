package main

import (
	routes "github.com/GorginZ/community-library/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/books", routes.HandleBooks)

	router.Run("localhost:8080")
}

// https://go.dev/doc/tutorial/web-service-gin
// tutorial

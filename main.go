package main

import (
	"log"

	metadata "github.com/GorginZ/community-library/metadata"
	routes "github.com/GorginZ/community-library/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/books", routes.HandleBooks)
	router.GET("/users", routes.HandleUsers)
	router.POST("/users", routes.AddUser)
	log.Printf("Version: %s", metadata.Version)

	log.Fatal(router.Run("0.0.0.0:8080"))
}

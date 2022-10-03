package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ISBN   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []book{
	{ISBN: "1", Title: "Brave New World", Author: "Aldous Huxley"},
	{ISBN: "2", Title: "Das Kapital Volume One", Author: "Karl Marx"},
	{ISBN: "3", Title: "If This Is A Man", Author: "Primo Levi"},
	{ISBN: "3", Title: "Sexus", Author: "Henry Miller"},
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// https://go.dev/doc/tutorial/web-service-gin
// tutorial

package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ISBN   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// instead of db
var (
	Books = []book{
		{ISBN: "1", Title: "Brave New World", Author: "Aldous Huxley"},
		{ISBN: "2", Title: "Das Kapital Volume One", Author: "Karl Marx"},
		{ISBN: "3", Title: "If This Is A Man", Author: "Primo Levi"},
		{ISBN: "3", Title: "Sexus", Author: "Henry Miller"},
	}
)

// getAlbums responds with the list of all albums as JSON.
func HandleBooks(c *gin.Context) {
	//sets content-type as application/json for us
	if Books != nil {
		c.JSON(http.StatusOK, Books)
		return
	}
	c.JSON(http.StatusInternalServerError, "unable to get books")
	return
}

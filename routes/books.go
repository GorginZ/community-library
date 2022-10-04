package routes

import (
	"net/http"

	br "github.com/GorginZ/community-library/repository"
	"github.com/gin-gonic/gin"
)

var (
	bookService br.BookService = br.NewBookService()
)

func HandleBooks(c *gin.Context) {
	Books, err := getBooks()
	//sets content-type as application/json for us
	if err != nil {
		c.JSON(http.StatusInternalServerError, "unable to get books")
		return
	}
	c.JSON(http.StatusOK, Books)
}

func getBooks() ([]br.Book, error) {
	return bookService.BookRepository.GetAll()
}

package routes

import (
	"net/http"

	br "github.com/GorginZ/community-library/repository"
	"github.com/gin-gonic/gin"
)

var (
	bookService *br.BookService = br.NewBookService()
)

func HandleBooks(c *gin.Context) {
	Books, err := getBooks()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(http.StatusOK, Books)
}

func getBooks() ([]br.Book, error) {
	books, err := bookService.BookRepository.GetAll()
	if err != nil {
		return books, err
	}
	return books, nil
}

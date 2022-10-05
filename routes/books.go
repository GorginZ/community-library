package routes

import (
	"net/http"

	br "github.com/GorginZ/community-library/repository"
	"github.com/gin-gonic/gin"
)

// how to give this the book service, is this bad? This should probably be on a struct
var (
	bookService *br.BookService = br.NewBookService()
)

func HandleBooks(c *gin.Context) {
	Books, err := getBooks()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	//sets content-type as application/json for us
	c.JSON(http.StatusOK, Books)
}

func getBooks() ([]br.Book, error) {
	//TODO err
	return bookService.BookRepository.GetAll()
}

// for testing this seems like a bad idea
func UseServiceWithFakeBookRepository() {
	bookService = br.NewBookService(br.WithFakeRepository())
}

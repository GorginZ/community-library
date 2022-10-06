package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	repository "github.com/GorginZ/community-library/repository"
	"github.com/gin-gonic/gin"
)

var bsWithEmptyRepo = repository.NewBookService(repository.WithFakeEmptyBookRepository())
var bsWithFullFakeRepo = repository.NewBookService(repository.WithFakeRepository())

func Test_Books(t *testing.T) {
	tests := map[string]struct {
		wantCode int
		request  *http.Request
		w        httptest.ResponseRecorder
		context  gin.Context
		bs       repository.BookService
	}{
		"get-books-should-return-200-and-books": {
			wantCode: 200,
			bs:       *bsWithFullFakeRepo,
		},
		// use 'real' client with no implementation to see sad path
		"get-books-should-return-500-see-sad-path": {
			wantCode: 500,
			bs:       *bsWithEmptyRepo,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := GetTestGinContext(tt.request, &tt.w)
			//set the bookservice
			bookService = &tt.bs

			HandleBooks(ctx)

			response := tt.w.Result()
			if response.StatusCode != tt.wantCode {
				t.Errorf("got %v, want %v", response.Status, tt.wantCode)
			}
		})
	}
}

func GetTestGinContext(r *http.Request, w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = r

	return ctx
}

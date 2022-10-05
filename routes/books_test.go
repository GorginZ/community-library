package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	bookRepository "github.com/GorginZ/community-library/repository"
	"github.com/gin-gonic/gin"
)

// fake bs
var bs *bookRepository.BookService = bookRepository.NewBookService(bookRepository.WithFakeRepository())

// func Test_getBooks(t *testing.T) {
// 	tests := map[string]struct {
// 		want    []bookRepository.Book
// 		wantErr error
// 	}{
// 		"getBooks-should-return-books": {
// 			//TODO  revisit this, it's testing nothing
// 			want: []bookRepository.Book{
// 				{ISBN: "1", Title: "Brave New World", Author: "Aldous Huxley"},
// 				{ISBN: "2", Title: "Das Kapital Volume One", Author: "Karl Marx"},
// 				{ISBN: "3", Title: "If This Is A Man", Author: "Primo Levi"},
// 				{ISBN: "3", Title: "Sexus", Author: "Henry Miller"},
// 			},
// 			wantErr: nil,
// 		},
// 	}
// 	for name, tt := range tests {
// 		t.Run(name, func(t *testing.T) {
// 			b, err := getBooks()
// 			if b == nil && err != tt.wantErr {
// 				t.Errorf("got %v, want %v", err, tt.want)
// 			}
// 		})
// 	}
// }

func Test_Books(t *testing.T) {
	tests := map[string]struct {
		wantCode int
		request  *http.Request
		w        httptest.ResponseRecorder
		context  gin.Context
	}{
		"get-books-should-return-200-and-books": {
			wantCode: 200,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {

			ctx := GetTestGinContext(tt.request, &tt.w)

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

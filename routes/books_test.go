package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	bookRepository "github.com/GorginZ/community-library/repository"
	"github.com/gin-gonic/gin"
)

func Test_getBooks(t *testing.T) {
	tests := map[string]struct {
		want    []bookRepository.Book
		wantErr error
	}{
		"getBooks-should-return-books": {
			//TODO refactor here and actually use the value, can set the var books
			want: []bookRepository.Book{
				{ISBN: "1", Title: "Brave New World", Author: "Aldous Huxley"},
				{ISBN: "2", Title: "Das Kapital Volume One", Author: "Karl Marx"},
				{ISBN: "3", Title: "If This Is A Man", Author: "Primo Levi"},
				{ISBN: "3", Title: "Sexus", Author: "Henry Miller"},
			},
			wantErr: nil,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			b, err := getBooks()
			if b == nil && err != tt.wantErr {
				t.Errorf("got %v, want %v", err, tt.want)
			}
		})
	}
}

func Test_Books(t *testing.T) {
	tests := map[string]struct {
		wantCode int
		request  *http.Request
		w        httptest.ResponseRecorder
		context  gin.Context
	}{
		"get-books-should-return-200-and-books": {
			wantCode: 200,
			context:  *GetTestGinContext(httptest.NewRequest("GET", "/books", nil)),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			HandleBooks(&tt.context)
			response := tt.w.Result()
			if response.StatusCode != tt.wantCode {
				t.Errorf("got %v, want %v", response.Status, tt.wantCode)
			}
		})
	}
}

func GetTestGinContext(r *http.Request) *gin.Context {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = r

	return ctx
}

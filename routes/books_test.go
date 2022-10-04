package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

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

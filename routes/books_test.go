package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// whole shebang kind of thing, service inc, so not really just testing books.go
func Test_Books(t *testing.T) {
	tests := map[string]struct {
		wantCode int
		request  *http.Request
		w        httptest.ResponseRecorder
		context  gin.Context
		useFake  bool
	}{
		"get-books-should-return-200-and-books": {
			wantCode: 200,
			useFake:  true,
		},
		//use 'real' client with no implementation to see sad path
		"get-books-should-return-500-see-sad-path": {
			wantCode: 500,
			useFake:  false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {

			ctx := GetTestGinContext(tt.request, &tt.w)
			if tt.useFake {
				UseServiceWithFakeBookRepository()
			}

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

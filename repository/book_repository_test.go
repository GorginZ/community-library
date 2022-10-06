package repository

import (
	"reflect"
	"testing"
)

var bsWithEmptyRepo *BookService = NewBookService(WithFakeEmptyBookRepository())
var bsWithFakeFullRepo *BookService = NewBookService(WithFakeRepository())

// should inject these
var books []Book = []Book{
	{ISBN: "1", Title: "Brave New World", Author: "Aldous Huxley"},
	{ISBN: "2", Title: "Das Kapital Volume One", Author: "Karl Marx"},
	{ISBN: "3", Title: "If This Is A Man", Author: "Primo Levi"},
	{ISBN: "3", Title: "Sexus", Author: "Henry Miller"},
}

// does the fake behave how I want, can other tests rely on it
func Test_Books(t *testing.T) {
	tests := map[string]struct {
		wantErr    bool
		wantResult []Book
		bs         *BookService
	}{
		"if-can't-get-books-should-err": {
			wantErr:    true,
			wantResult: nil,
			bs:         bsWithEmptyRepo,
		},
		"if-can-get-books-should-be-no-err": {
			wantErr:    false,
			wantResult: books,
			bs:         bsWithFakeFullRepo,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := tt.bs.BookRepository.GetAll()
			var gotErr bool
			if err != nil {
				gotErr = true
			}

			if !reflect.DeepEqual(result, tt.wantResult) || gotErr != tt.wantErr {
				t.Errorf("want books: %v got books: %v, want error %v, got error %v", tt.wantResult, result, tt.wantErr, gotErr)
			}
		})
	}
}

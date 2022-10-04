package repository

import (
	"testing"
)

func Test_opts(t *testing.T) {
	bs := NewBookService(WithFakeRepository())
	result, err := bs.BookRepository.GetAll()

	if result == nil && err != nil {
		t.Errorf("fake not return books, got %v, want %v", result, err)
	}
}

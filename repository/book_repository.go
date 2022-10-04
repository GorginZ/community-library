package repository

import "errors"

type Book struct {
	ISBN   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// the book repository iface
type BookGetter interface {
	GetAll() ([]Book, error)
}

// the service layer
type BookService struct {
	//book repository can prob be private idk. Maybe that's too many layers though
	BookRepository BookGetter
}

// new default service
func NewBookService(opts ...BookServiceOption) *BookService {
	// defaults
	defaultRepository := BookRepository{}
	bs := &BookService{
		BookRepository: defaultRepository,
	}
	for _, opt := range opts {
		opt(bs)
	}
	return bs
}

type BookRepository struct {
}

type BookServiceOption func(*BookService)

// FakeBookRepository BookServiceOption for testing
func WithFakeRepository() BookServiceOption {
	return func(bs *BookService) {
		bs.BookRepository = newFakeBookRepository()
	}
}

func (r BookRepository) GetAll() ([]Book, error) {
	return nil, errors.New("not implemented")
}

// book repository is a BookGetter
type FakeBookRepository struct {
	Books []Book
}

func newFakeBookRepository() FakeBookRepository {
	br := FakeBookRepository{Books: []Book{
		{ISBN: "1", Title: "Brave New World", Author: "Aldous Huxley"},
		{ISBN: "2", Title: "Das Kapital Volume One", Author: "Karl Marx"},
		{ISBN: "3", Title: "If This Is A Man", Author: "Primo Levi"},
		{ISBN: "3", Title: "Sexus", Author: "Henry Miller"},
	}}
	return br
}

func (r FakeBookRepository) GetAll() ([]Book, error) {
	return r.Books, nil
}

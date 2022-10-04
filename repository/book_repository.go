package repository

type Book struct {
	ISBN   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// instead of db
var (
	Books = []Book{
		{ISBN: "1", Title: "Brave New World", Author: "Aldous Huxley"},
		{ISBN: "2", Title: "Das Kapital Volume One", Author: "Karl Marx"},
		{ISBN: "3", Title: "If This Is A Man", Author: "Primo Levi"},
		{ISBN: "3", Title: "Sexus", Author: "Henry Miller"},
	}
)

// the book repository iface
type BookGetter interface {
	GetAll() ([]Book, error)
}

type BookService struct {
	BookRepository BookGetter
}

func NewBookService() BookService {
	r := NewFakeBookRepository()
	s := BookService{BookRepository: r}
	return s
}

type FakeBookRepository struct {
	Books []Book
}

func NewFakeBookRepository() FakeBookRepository {
	br := FakeBookRepository{Books: []Book{
		{ISBN: "1", Title: "Brave New World", Author: "Aldous Huxley"},
		{ISBN: "2", Title: "Das Kapital Volume One", Author: "Karl Marx"},
		{ISBN: "3", Title: "If This Is A Man", Author: "Primo Levi"},
		{ISBN: "3", Title: "Sexus", Author: "Henry Miller"},
	}}
	return br
}

func (r FakeBookRepository) GetAll() ([]Book, error) {
	return Books, nil
}

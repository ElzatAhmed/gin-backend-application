package model

type GenreType string

var bookSequence uint64 = 0

const (
	Fiction = "fiction"
	Mystery = "mystery"
	Horror = "horror"
	Thriller = "thriller"
	Historical = "historical"
	Romance = "romance"
	Western = "western"
)

type Book struct {
	Id			uint64		`json:"id"`
	Name		string		`json:"name"`
	Genre   	GenreType	`json:"genre"`
	AuthorId 	uint64		`json:"author_id"`
}

func NewBook(name string, genre GenreType, authorId uint64) *Book {
	book := Book{
		Id: 	  bookSequence,
		Name:     name,
		Genre:    genre,
		AuthorId: authorId,
	}
	bookSequence += 1
	return &book
}

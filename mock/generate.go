package mock

import (
	"fmt"
	"github.com/gin-backend-application/model"
	"math/rand"
)

func GenerateMockData(count int) ([]*model.Author, []*model.Book) {
	genre := [7]model.GenreType{model.Fiction, model.Mystery, model.Historical, model.Horror,
		model.Romance, model.Thriller, model.Western}
	var authors [] *model.Author
	var books []*model.Book
	for i := 0; i < count; i++ {
		author := model.NewAuthor(fmt.Sprintf("author_%d", i),
			fmt.Sprintf("nation_%d", i))
		authors = append(authors, author)
		for j := 0; j < count % 10; j++ {
			book := model.NewBook(
				fmt.Sprintf("book_%d_%d", author.Id, j),
				genre[rand.Intn(6)],
				author.Id)
			books = append(books, book)
		}
	}
	return authors, books
}

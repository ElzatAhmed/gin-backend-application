package author_service_implements

import (
	"github.com/gin-backend-application/model"
	"github.com/gin-backend-application/repository"
	"github.com/gin-backend-application/service"
)

type authorPublishServiceImpl struct {
	authorRepo repository.AuthorRepository
	bookRepo repository.BookRepository
}

func NewAuthorPublishService(authorRepo repository.AuthorRepository,
	bookRepo repository.BookRepository) service.AuthorPublishService {
	return &authorPublishServiceImpl{
		authorRepo: authorRepo,
		bookRepo:   bookRepo,
	}
}

func (a *authorPublishServiceImpl) PublishNewBook(authorId uint64,
	bookInfo model.Book) (model.Book, error) {
	_, err := a.authorRepo.FindById(authorId)
	if err != nil {
		return model.Book{}, err
	}
	newBook, err := a.bookRepo.SaveNew(bookInfo)
	if err != nil {
		return model.Book{}, err
	}
	return newBook, nil
}
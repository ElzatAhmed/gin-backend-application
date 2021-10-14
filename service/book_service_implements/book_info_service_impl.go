package book_service_implements

import (
	"errors"
	"fmt"
	"github.com/gin-backend-application/model"
	"github.com/gin-backend-application/repository"
	"github.com/gin-backend-application/service"
)

type bookInfoServiceImpl struct {
	bookRepo repository.BookRepository
	authorRepo repository.AuthorRepository
}

func NewBookInfoService(authorRepo repository.AuthorRepository,
	bookRepo repository.BookRepository) service.BookInfoService {
	return &bookInfoServiceImpl{
		bookRepo:   bookRepo,
		authorRepo: authorRepo,
	}
}

func (b *bookInfoServiceImpl) FindById(id uint64) (model.Book, error) {
	return b.bookRepo.FindById(id)
}

func (b *bookInfoServiceImpl) FindAllByName(name string) ([]model.Book, error) {
	return b.bookRepo.FindAllByName(name)
}

func (b *bookInfoServiceImpl) FindAllByAuthorId(authorId uint64) ([]model.Book, error) {
	return b.bookRepo.FindAllByAuthorId(authorId)
}

func (b bookInfoServiceImpl) FindAllByAuthorName(authorName string) ([]model.Book, error) {
	authors, err := b.authorRepo.FindAllByName(authorName)
	if err != nil {
		return nil, err
	}
	var res []model.Book
	for _, author := range authors{
		books, err := b.bookRepo.FindAllByAuthorId(author.Id)
		if err != nil {
			continue
		}
		for _, book := range books{
			res = append(res, book)
		}
	}
	if len(res) == 0{
		return nil, errors.New(fmt.Sprintf("books with author name %s not found", authorName))
	}
	return res, nil
}


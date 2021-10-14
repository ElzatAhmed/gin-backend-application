package repository

import (
	"errors"
	"fmt"
	"github.com/gin-backend-application/model"
)

type BookRepository interface {
	FindById(id uint64) (model.Book, error)
	FindAllByName(name string) ([]model.Book,error)
	UpdateById(id uint64, newInfo model.Book) (model.Book,error)
	SaveNew(info model.Book) (model.Book,error)
}

type bookRepoImpl struct {
	books []*model.Book
}

func (b *bookRepoImpl) FindById(id uint64) (model.Book,error) {
	book, err := b.findBookById(id)
	if err != nil {
		return model.Book{}, err
	}
	return *book, nil
}

func (b *bookRepoImpl) FindAllByName(name string) ([]model.Book,error) {
	var books []model.Book
	for _, book := range b.books{
		if book.Name == name{
			books = append(books, *book)
		}
	}
	if len(books) ==0{
		return books, errors.New(fmt.Sprintf("book with name %s not found", name))
	}
	return books, nil
}

func (b *bookRepoImpl) UpdateById(id uint64, newInfo model.Book) (model.Book,error) {
	book, err := b.findBookById(id)
	if err != nil {
		return model.Book{}, err
	}
	book.Name = newInfo.Name
	book.AuthorId = newInfo.AuthorId
	book.Genre = newInfo.Genre
	return *book, nil
}

func (b *bookRepoImpl) SaveNew(info model.Book) (model.Book,error) {
	newBook := model.NewBook(info.Name, info.Genre, info.AuthorId)
	b.books = append(b.books, newBook)
	return *newBook, nil
}


// private methods
func (b *bookRepoImpl) findBookById(id uint64) (*model.Book, error) {
	for _, book := range b.books{
		if book.Id == id{
			return book, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("book with id %d not found", id))
}


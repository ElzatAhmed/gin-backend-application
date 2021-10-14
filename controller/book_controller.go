package controller

import "github.com/gin-backend-application/model"

type BookInfoController interface {
	GetBookById(id uint64) (model.Book, error)
	GetBooksByName(name string) ([]model.Book, error)
	GetBooksByAuthorId(id uint64) ([]model.Book, error)
	GetBooksByAuthorName(name string) ([]model.Book, error)
}

type BookUpdateController interface {
	UpdateBookInfoById(id uint64, book model.Book) (model.Book, error)
}
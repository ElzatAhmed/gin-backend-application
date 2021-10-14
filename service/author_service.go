package service

import (
	"github.com/gin-backend-application/model"
)


type AuthorInfoService interface {
	FindById(id uint64) (model.Author, error)
	FindAllByName(name string) ([]model.Author,error)
	UpdateInfoById(id uint64, newInfo model.Author) (model.Author,error)
}

type AuthorPublishingService interface {
	PublishNewBook(authorId uint64, bookInfo model.Book) (model.Book, error)
}

type AuthorRegisteringService interface {
	RegisterNewAuthor(authorInfo model.Author) (model.Author, error)
}


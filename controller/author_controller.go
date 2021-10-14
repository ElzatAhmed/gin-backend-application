package controller

import "github.com/gin-backend-application/model"

type AuthorInfoController interface {
	GetAuthorById(id uint64) (model.Author, error)
	GetAuthorsByName(name string) ([]model.Author, error)
	UpdateAuthorInfoById(id uint64, author model.Author) (model.Author, error)
}

type AuthorPublishController interface {
	Publish(authorId uint64, book model.Book) (model.Book, error)
}

type AuthorRegisterController interface {
	Register(author model.Author) (model.Author, error)
}

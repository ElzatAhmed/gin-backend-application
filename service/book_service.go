package service

import "github.com/gin-backend-application/model"

type BookInfoService interface {
	FindById(id uint64) (model.Book, error)
	FindAllByName(name string) ([]model.Book, error)
	FindAllByAuthorId(authorId uint64) ([]model.Book, error)
	FindAllByAuthorName(authorName string) ([]model.Book, error)
}

type BookUpdatingService interface {
	UpdateById(id uint64, newInfo model.Book) (model.Book, error)
}
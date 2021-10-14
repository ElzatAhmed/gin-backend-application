package repository

import (
	"github.com/gin-backend-application/model"
)

type BookRepository interface {
	FindById(id uint64) (model.Book, error)
	FindAllByName(name string) ([]model.Book,error)
	FindAllByAuthorId(id uint64) ([]model.Book, error)
	UpdateById(id uint64, newInfo model.Book) (model.Book,error)
	SaveNew(info model.Book) (model.Book,error)
}


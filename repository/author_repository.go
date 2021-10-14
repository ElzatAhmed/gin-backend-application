package repository

import (
	"github.com/gin-backend-application/model"
)

type AuthorRepository interface {
	FindById(id uint64) (model.Author, error)
	FindAllByName(name string) ([]model.Author, error)
	UpdateById(id uint64, newInfo model.Author) (model.Author, error)
	SaveNew(info model.Author) (model.Author, error)
}



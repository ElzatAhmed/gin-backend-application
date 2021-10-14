package book_service_implements

import (
	"github.com/gin-backend-application/model"
	"github.com/gin-backend-application/repository"
	"github.com/gin-backend-application/repository/repo_implements"
	"github.com/gin-backend-application/service"
)

type bookUpdatingServiceImpl struct {
	bookRepo repository.BookRepository
}

func NewBookUpdatingService() service.BookUpdatingService {
	return &bookUpdatingServiceImpl{
		bookRepo: repo_implements.NewBookRepo(),
	}
}

func (b bookUpdatingServiceImpl) UpdateById(id uint64, newInfo model.Book) (model.Book, error) {
	return b.bookRepo.UpdateById(id, newInfo)
}


package author_service_implements

import (
	"github.com/gin-backend-application/model"
	"github.com/gin-backend-application/repository"
	"github.com/gin-backend-application/service"
)

type authorInfoServiceImpl struct {
	authorRepo repository.AuthorRepository
}

func NewAuthorInfoService(authorRepo repository.AuthorRepository) service.AuthorInfoService {
	return &authorInfoServiceImpl{
		authorRepo: authorRepo,
	}
}

func (a *authorInfoServiceImpl) FindById(id uint64) (model.Author, error) {
	return a.authorRepo.FindById(id)
}

func (a *authorInfoServiceImpl) FindAllByName(name string) ([]model.Author, error) {
	return a.authorRepo.FindAllByName(name)
}

func (a *authorInfoServiceImpl) UpdateInfoById(id uint64, newInfo model.Author) (model.Author, error) {
	return a.authorRepo.UpdateById(id, newInfo)
}
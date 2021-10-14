package author_service_implements

import (
	"github.com/gin-backend-application/model"
	"github.com/gin-backend-application/repository"
	"github.com/gin-backend-application/service"
)

type authorRegisteringServiceImpl struct {
	authorRepo repository.AuthorRepository
}

func NewAuthorRegisteringService(authorRepo repository.AuthorRepository) service.AuthorRegisteringService {
	return &authorRegisteringServiceImpl{
		authorRepo: authorRepo,
	}
}

func (a *authorRegisteringServiceImpl) RegisterNewAuthor(authorInfo model.Author) (model.Author, error) {
	author, err := a.authorRepo.SaveNew(authorInfo)
	if err != nil {
		return model.Author{}, err
	}
	return author, nil
}




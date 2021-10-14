package author_service_implements

import (
	"github.com/gin-backend-application/model"
	"github.com/gin-backend-application/repository"
	"github.com/gin-backend-application/service"
)

type authorRegisterServiceImpl struct {
	authorRepo repository.AuthorRepository
}

func NewAuthorRegisterService(authorRepo repository.AuthorRepository) service.AuthorRegisterService {
	return &authorRegisterServiceImpl{
		authorRepo: authorRepo,
	}
}

func (a *authorRegisterServiceImpl) RegisterNewAuthor(authorInfo model.Author) (model.Author, error) {
	author, err := a.authorRepo.SaveNew(authorInfo)
	if err != nil {
		return model.Author{}, err
	}
	return author, nil
}




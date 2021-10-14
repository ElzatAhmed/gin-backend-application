package author_controller_implements

import (
	"github.com/gin-backend-application/controller"
	"github.com/gin-backend-application/model"
	"github.com/gin-backend-application/service"
)

type authorInfoControllerImpl struct {
	authorInfoService service.AuthorInfoService
}

func NewAuthorInfoController(infoService service.AuthorInfoService) controller.AuthorInfoController {
	return &authorInfoControllerImpl{
		authorInfoService: infoService,
	}
}

func (a *authorInfoControllerImpl) GetAuthorById(id uint64) (model.Author, error) {
	return a.authorInfoService.FindById(id)
}

func (a *authorInfoControllerImpl) GetAuthorsByName(name string) ([]model.Author, error) {
	return a.authorInfoService.FindAllByName(name)
}

func (a *authorInfoControllerImpl) UpdateAuthorInfoById(id uint64,
	author model.Author) (model.Author, error) {
	return a.authorInfoService.UpdateInfoById(id, author)
}


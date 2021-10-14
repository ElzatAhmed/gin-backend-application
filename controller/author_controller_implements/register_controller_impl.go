package author_controller_implements

import (
	"github.com/gin-backend-application/controller"
	"github.com/gin-backend-application/model"
	"github.com/gin-backend-application/service"
)

type authorRegisterServiceImpl struct {
	authorRegisterService service.AuthorRegisterService
}

func NewAuthorRegisterController(registerService service.AuthorRegisterService) controller.AuthorRegisterController {
	return &authorRegisterServiceImpl{
		authorRegisterService: registerService,
	}
}

func (a *authorRegisterServiceImpl) Register(author model.Author) (model.Author, error) {
	return a.authorRegisterService.RegisterNewAuthor(author)
}


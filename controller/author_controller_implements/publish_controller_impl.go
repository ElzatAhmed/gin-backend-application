package author_controller_implements

import (
	"github.com/gin-backend-application/controller"
	"github.com/gin-backend-application/model"
	"github.com/gin-backend-application/service"
)

type authorPublishControllerImpl struct {
	authorPublishService service.AuthorPublishService
}

func NewAuthorPublishController(publishService service.AuthorPublishService) controller.AuthorPublishController {
	return &authorPublishControllerImpl{
		authorPublishService: publishService,
	}
}

func (a *authorPublishControllerImpl) Publish(authorId uint64, book model.Book) (model.Book, error) {
	return a.authorPublishService.PublishNewBook(authorId, book)
}



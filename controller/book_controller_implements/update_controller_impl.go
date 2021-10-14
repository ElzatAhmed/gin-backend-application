package book_controller_implements

import (
	"github.com/gin-backend-application/controller"
	"github.com/gin-backend-application/model"
	"github.com/gin-backend-application/service"
)

type bookUpdateControllerImpl struct {
	bookUpdateService service.BookUpdateService
}

func NewBookUpdateController(updateService service.BookUpdateService) controller.BookUpdateController {
	return &bookUpdateControllerImpl{
		bookUpdateService: updateService,
	}
}

func (b *bookUpdateControllerImpl) UpdateBookInfoById(id uint64, book model.Book) (model.Book, error) {
	return b.bookUpdateService.UpdateById(id, book)
}


package book_controller_implements

import (
	"github.com/gin-backend-application/controller"
	"github.com/gin-backend-application/model"
	"github.com/gin-backend-application/service"
)

type bookInfoControllerImpl struct {
	bookInfoService service.BookInfoService
}

func NewBookInfoController(infoService service.BookInfoService) controller.BookInfoController {
	return &bookInfoControllerImpl{
		bookInfoService: infoService,
	}
}

func (b *bookInfoControllerImpl) GetBookById(id uint64) (model.Book, error) {
	return b.bookInfoService.FindById(id)
}

func (b *bookInfoControllerImpl) GetBooksByName(name string) ([]model.Book, error) {
	return b.GetBooksByName(name)
}

func (b *bookInfoControllerImpl) GetBooksByAuthorId(id uint64) ([]model.Book, error) {
	return b.bookInfoService.FindAllByAuthorId(id)
}

func (b *bookInfoControllerImpl) GetBooksByAuthorName(name string) ([]model.Book, error) {
	return b.bookInfoService.FindAllByAuthorName(name)
}




package repository

import (
	"errors"
	"fmt"
	"github.com/gin-backend-application/model"
)

type AuthorRepository interface {
	FindById(id uint64) (model.Author, error)
	FindAllByName(name string) ([]model.Author, error)
	UpdateById(id uint64, newInfo model.Author) (model.Author, error)
	SaveNew(info model.Author) (model.Author, error)
}

// implementation
type authorRepoImpl struct {
	authors []*model.Author	// simulates the database
}

func NewAuthorRepo() AuthorRepository{
	return &authorRepoImpl{}
}

func (a *authorRepoImpl) FindById(id uint64) (model.Author, error) {
	author, err := a.findAuthorById(id)
	if err != nil{
		return model.Author{}, err
	}
	return *author, nil
}

func (a *authorRepoImpl) FindAllByName(name string) ([]model.Author, error) {
	var authors []model.Author
	for _, author := range a.authors{
		if author.Name == name{
			authors = append(authors, *author)
		}
	}
	if len(authors) == 0{
		return authors, errors.New(fmt.Sprintf("author with name %s not found", name))
	}
	return authors, nil
}

func (a *authorRepoImpl) UpdateById(id uint64, newInfo model.Author) (model.Author, error) {
	author, err := a.findAuthorById(id)
	if err != nil{
		return model.Author{}, err
	}
	author.Name = newInfo.Name
	author.Nationality = newInfo.Nationality
	return *author, nil
}

func (a *authorRepoImpl) SaveNew(info model.Author) (model.Author, error) {
	author := model.NewAuthor(info.Name, info.Nationality)
	a.authors = append(a.authors, author)
	return *author, nil
}


// private methods
func (a *authorRepoImpl) findAuthorById(id uint64) (*model.Author, error){
	for _, author := range a.authors{
		if author.Id == id{
			return author, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("author with id %d not found", id))
}

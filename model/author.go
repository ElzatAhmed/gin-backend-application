package model

var authorSequence uint64 = 0

type Author struct {
	Id          uint64 `json:"Id"`
	Name        string `json:"Name"`
	Nationality string `json:"nationality"`
}

func NewAuthor(name string, nationality string) *Author {
	author := &Author{Id: authorSequence, Name: name, Nationality: nationality}
	authorSequence += 1
	return author
}
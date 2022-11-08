package books

import (
	"errors"

	"gorm.io/gorm"
)

type book_repo struct {
	db *gorm.DB
}

func NewRepo(grm *gorm.DB) *book_repo {
	return &book_repo{grm}
}

func (r *book_repo) FindAll() (*Books, error) {
	var books Books

	result := r.db.Find(&books)

	if result.Error != nil {
		return nil, errors.New("Retrieve data failure")
	}

	return &books, nil
}

func (r *book_repo) Add(data *Book) (*Book, error) {

	result := r.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("Add data failure")
	}

	return data, nil
}

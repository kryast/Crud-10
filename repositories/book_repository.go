package repositories

import (
	"github.com/kryast/Crud-10.git/models"
	"gorm.io/gorm"
)

type BookRepository interface {
	Create(book *models.Book) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepistory(db *gorm.DB) BookRepository {
	return &bookRepository{db}
}

func (br *bookRepository) Create(book *models.Book) error {
	return br.db.Create(book).Error
}

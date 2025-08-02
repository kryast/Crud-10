package repositories

import (
	"github.com/kryast/Crud-10.git/models"
	"gorm.io/gorm"
)

type BookRepository interface {
	Create(book *models.Book) error
	GetAll() ([]models.Book, error)
	GetByID(id uint) (*models.Book, error)
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

func (br *bookRepository) GetAll() ([]models.Book, error) {
	var books []models.Book
	err := br.db.Find(&books).Error

	return books, err
}

func (br *bookRepository) GetByID(id uint) (*models.Book, error) {
	var book models.Book
	err := br.db.First(&book, id).Error

	return &book, err
}

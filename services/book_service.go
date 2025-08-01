package services

import "github.com/kryast/Crud-10.git/repositories"

type BookService interface {
}

type bookService struct {
	repo repositories.BookRepository
}

func NewBookService(repo repositories.BookRepository) BookService {
	return &bookService{repo}
}

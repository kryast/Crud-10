package handlers

import "github.com/kryast/Crud-10.git/services"

type BookHandler struct {
	bs services.BookService
}

func NewBookHandler(bs services.BookService) *BookHandler {
	return &BookHandler{bs}
}

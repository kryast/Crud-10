package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kryast/Crud-10.git/models"
	"github.com/kryast/Crud-10.git/services"
)

type BookHandler struct {
	bs services.BookService
}

func NewBookHandler(bs services.BookService) *BookHandler {
	return &BookHandler{bs}
}

func (bh *BookHandler) Create(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err})
	}
	if err := bh.bs.Create(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
	}

	c.JSON(http.StatusCreated, book)
}

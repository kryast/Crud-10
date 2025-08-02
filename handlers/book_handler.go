package handlers

import (
	"net/http"
	"strconv"

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
		return
	}
	if err := bh.bs.Create(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}

	c.JSON(http.StatusCreated, book)
}

func (bh *BookHandler) GetAll(c *gin.Context) {
	books, err := bh.bs.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}
	c.JSON(http.StatusOK, books)

}

func (bh *BookHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	book, err := bh.bs.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}
	c.JSON(http.StatusOK, book)
}

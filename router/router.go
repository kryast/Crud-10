package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kryast/Crud-10.git/handlers"
	"github.com/kryast/Crud-10.git/repositories"
	"github.com/kryast/Crud-10.git/services"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	repo := repositories.NewBookRepistory(db)
	svc := services.NewBookService(repo)
	h := handlers.NewBookHandler(svc)

	r.POST("/book", h.Create)
	r.GET("/book", h.GetAll)
	r.GET("/book/:id", h.GetByID)
	r.PUT("/book/:id", h.Update)

	return r
}

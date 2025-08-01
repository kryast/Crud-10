package main

import (
	"log"

	"github.com/kryast/Crud-10.git/database"
	"github.com/kryast/Crud-10.git/models"
	"github.com/kryast/Crud-10.git/router"
)

func Main() {

	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to Connect Database", err)
	}

	db.AutoMigrate(&models.Book{})

	r := router.SetupRouter(db)
	r.Run(":8080")
}

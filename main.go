package main

import (
	"book-rental/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/books", handlers.AddBook)
	r.GET("/books", handlers.ListBooks)
	r.POST("/users", handlers.CreateUser)
	r.POST("/rent", handlers.RentBook)
	r.POST("/return", handlers.ReturnBook)

	r.Run(":7000")
}

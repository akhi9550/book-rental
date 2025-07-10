package handlers

import (
	"book-rental/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var Books = make(map[string]models.Book)
var Users = make(map[string]models.User)

func AddBook(c *gin.Context) {
	var input struct {
		Title  string `json:"title" binding:"required"`
		Author string `json:"author" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{
		ID:        uuid.New(),
		Title:     input.Title,
		Author:    input.Author,
		Available: true,
	}
	Books[book.ID.String()] = book
	c.JSON(http.StatusCreated, book)
}

func ListBooks(c *gin.Context) {
	var books []models.Book
	for _, book := range Books {
		books = append(books, book)
	}
	c.JSON(http.StatusOK, books)
}

func CreateUser(c *gin.Context) {
	var input struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, u := range Users {
		if u.Email == input.Email {
			c.JSON(http.StatusConflict, gin.H{"error": "email already exists"})
			return
		}
	}

	user := models.User{
		ID:    uuid.New(),
		Name:  input.Name,
		Email: input.Email,
	}

	Users[user.ID.String()] = user
	c.JSON(http.StatusCreated, user)
}

func RentBook(c *gin.Context) {
	var input struct {
		UserID string `json:"user_id" binding:"required"`
		BookID string `json:"book_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, ok := Books[input.BookID]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}

	if !book.Available {
		c.JSON(http.StatusConflict, gin.H{"error": "book already rented"})
		return
	}

	if _, exists := Users[input.UserID]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	book.Available = false
	Books[input.BookID] = book
	c.JSON(http.StatusOK, gin.H{"message": "book rented successfully"})
}

func ReturnBook(c *gin.Context) {
	var input struct {
		BookID string `json:"book_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, ok := Books[input.BookID]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}

	book.Available = true
	Books[input.BookID] = book
	c.JSON(http.StatusOK, gin.H{"message": "book returned successfully"})
}

package tests

import (
	"book-rental/handlers"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAddBook(t *testing.T) {
	router := gin.Default()
	router.POST("/books", handlers.AddBook)

	payload := map[string]string{"title": "The Psychology of Money", "author": "Morgan Housel"}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, 201, resp.Code)
}

func TestCreateUser(t *testing.T) {
	router := gin.Default()
	router.POST("/users", handlers.CreateUser)

	payload := map[string]string{"name": "abhilash", "email": "abhilash@example.com"}
	body, _ := json.Marshal(payload)

	req1, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	req1.Header.Set("Content-Type", "application/json")
	resp1 := httptest.NewRecorder()
	router.ServeHTTP(resp1, req1)
	assert.Equal(t, http.StatusCreated, resp1.Code)

	req2, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	req2.Header.Set("Content-Type", "application/json")
	resp2 := httptest.NewRecorder()
	router.ServeHTTP(resp2, req2)
	assert.Equal(t, http.StatusConflict, resp2.Code)
}

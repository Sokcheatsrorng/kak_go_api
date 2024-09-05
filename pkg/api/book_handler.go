package api

import (
	"author_book_api/pkg/models"
	"author_book_api/pkg/service"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

// GetAllBooks retrieves all Books.
func GetAllBooks(BookService *service.BookService) gin.HandlerFunc {
    return func(c *gin.Context) {
        Books, err := BookService.GetAllBooks()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, Books)
    }
}

// GetBook retrieves an Book by ID.
func GetBook(BookService *service.BookService) gin.HandlerFunc {
    return func(c *gin.Context) {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
            return
        }
        Book, err := BookService.GetBookByID(uint(id))
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, Book)
    }
}

// CreateBook creates a new Book.
func CreateBook(BookService *service.BookService) gin.HandlerFunc {
    return func(c *gin.Context) {
        var Book models.Book
        if err := c.BindJSON(&Book); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
            return
        }
        createdBook, err := BookService.CreateBook(Book)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusCreated, createdBook)
    }
}

// UpdateBook updates an existing Book.
func UpdateBook(BookService *service.BookService) gin.HandlerFunc {
    return func(c *gin.Context) {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
            return
        }
        var Book models.Book
        if err := c.BindJSON(&Book); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
            return
        }
        updatedBook, err := BookService.UpdateBook(uint(id), Book)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, updatedBook)
    }
}

// DeleteBook deletes an Book by ID.
func DeleteBook(BookService *service.BookService) gin.HandlerFunc {
    return func(c *gin.Context) {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
            return
        }
        err = BookService.DeleteBook(uint(id))
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.Status(http.StatusNoContent)
    }
}

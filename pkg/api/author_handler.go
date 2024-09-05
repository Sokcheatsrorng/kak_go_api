package api

import (
	"author_book_api/pkg/models"
	"author_book_api/pkg/service"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

// GetAllAuthors retrieves all authors.
func GetAllAuthors(authorService *service.AuthorService) gin.HandlerFunc {
    return func(c *gin.Context) {
        authors, err := authorService.GetAllAuthors()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, authors)
    }
}

// GetAuthor retrieves an author by ID.
func GetAuthor(authorService *service.AuthorService) gin.HandlerFunc {
    return func(c *gin.Context) {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
            return
        }
        author, err := authorService.GetAuthorByID(uint(id))
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, author)
    }
}

// CreateAuthor creates a new author.
func CreateAuthor(authorService *service.AuthorService) gin.HandlerFunc {
    return func(c *gin.Context) {
        var author models.Author
        if err := c.BindJSON(&author); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
            return
        }
        createdAuthor, err := authorService.CreateAuthor(author)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusCreated, createdAuthor)
    }
}

// UpdateAuthor updates an existing author.
func UpdateAuthor(authorService *service.AuthorService) gin.HandlerFunc {
    return func(c *gin.Context) {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
            return
        }
        var author models.Author
        if err := c.BindJSON(&author); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
            return
        }
        updatedAuthor, err := authorService.UpdateAuthor(uint(id), author)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, updatedAuthor)
    }
}

// DeleteAuthor deletes an author by ID.
func DeleteAuthor(authorService *service.AuthorService) gin.HandlerFunc {
    return func(c *gin.Context) {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
            return
        }
        err = authorService.DeleteAuthor(uint(id))
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.Status(http.StatusNoContent)
    }
}

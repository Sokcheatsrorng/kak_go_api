package api

import (
	"author_book_api/pkg/service"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, authorService *service.AuthorService, bookService *service.BookService) {
    // Author routes
    authorGroup := router.Group("/authors")
    {
        authorGroup.GET("", GetAllAuthors(authorService))
        authorGroup.GET("/:id", GetAuthor(authorService))
        authorGroup.POST("", CreateAuthor(authorService)) 
        authorGroup.PUT("/:id", UpdateAuthor(authorService)) 
        authorGroup.DELETE("/:id", DeleteAuthor(authorService)) 
	
    }

    // Book routes
    bookGroup := router.Group("/books")
    {
        bookGroup.GET("", GetAllBooks(bookService))
        bookGroup.GET("/:id", GetBook(bookService))
        bookGroup.POST("", CreateBook(bookService)) 
        bookGroup.PUT("/:id", UpdateBook(bookService)) 
        bookGroup.DELETE("/:id", DeleteBook(bookService)) 
    }
}

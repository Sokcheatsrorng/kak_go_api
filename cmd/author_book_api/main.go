package main

import (
	"author_book_api/pkg/api"
	"author_book_api/pkg/configs"
	"author_book_api/pkg/database"
	"author_book_api/pkg/repository"
	"author_book_api/pkg/service"
	"log"

	"github.com/gin-gonic/gin"
	// "github.com/swaggo/gin-swagger"
	// "github.com/swaggo/gin-swagger/swaggerFiles"
	//  _ "author_book_api/docs"
)


   

func main() {
    // Load configuration
    cfg := configs.LoadConfig()

    // Initialize database
    db := database.InitDB(cfg.DatabaseDSN)

    // Initialize repositories and services
    authorRepo := repository.NewAuthorRepo(db)
    bookRepo := repository.NewBookRepo(db)
    authorService := service.NewAuthorService(authorRepo)
    bookService := service.NewBookService(bookRepo)

    // Initialize Gin router
    router := gin.Default()
    api.SetupRouter(router, authorService, bookService)

	 // Swagger route
	//  router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


    // Start server
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("failed to start server: %v", err)
    }
}

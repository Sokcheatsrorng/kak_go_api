package database

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
    "author_book_api/pkg/models"
)

// InitDB initializes the SQLite database and performs automigration.
func InitDB(dsn string) *gorm.DB {

    db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }

    // Perform automatic migration
    err = db.AutoMigrate(&models.Author{}, &models.Book{})
    if err != nil {
        log.Fatalf("failed to migrate database: %v", err)
    }

    return db;
}
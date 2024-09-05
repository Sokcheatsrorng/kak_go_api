package repository

import (
	"author_book_api/pkg/models"
	"gorm.io/gorm"
)

type BookRepo struct {
    DB *gorm.DB
}

func NewBookRepo(db *gorm.DB) *BookRepo {
    return &BookRepo{DB: db}
}

func (repo *BookRepo) GetAll() ([]models.Book, error) {
    var books []models.Book
    err := repo.DB.Find(&books).Error
    return books, err
}

func (repo *BookRepo) GetByID(id uint) (models.Book, error) {
    var book models.Book
    err := repo.DB.First(&book, id).Error
    return book, err
}

func (repo *BookRepo) Create(book models.Book) (models.Book, error) {
    err := repo.DB.Create(&book).Error
    return book, err
}

func (repo *BookRepo) Update(id uint, book models.Book) (models.Book, error) {
    var existingBook models.Book
    if err := repo.DB.First(&existingBook, id).Error; err != nil {
        return existingBook, err
    }
    existingBook = book
    existingBook.ID = id
    err := repo.DB.Save(&existingBook).Error
    return existingBook, err
}

func (repo *BookRepo) Delete(id uint) error {
    return repo.DB.Delete(&models.Book{}, id).Error
}

package repository

import (
	"author_book_api/pkg/models"
	"gorm.io/gorm"
)

type AuthorRepo struct {
    DB *gorm.DB
}

func NewAuthorRepo(db *gorm.DB) *AuthorRepo {
    return &AuthorRepo{DB: db}
}

func (repo *AuthorRepo) GetAllAuthors() ([]models.Author, error) {
    var authors []models.Author
    err := repo.DB.Preload("Books").Find(&authors).Error
    return authors, err
}

func (repo *AuthorRepo) GetByID(id uint) (models.Author, error) {
    var author models.Author
    err := repo.DB.Preload("Books").First(&author, id).Error
    return author, err
}

func (repo *AuthorRepo) CreateAuthor(author models.Author) (models.Author, error) {
    err := repo.DB.Create(&author).Error
    return author, err
}

func (repo *AuthorRepo) UpdateAuthor(id uint, author models.Author) (models.Author, error) {
    var existingAuthor models.Author
    if err := repo.DB.First(&existingAuthor, id).Error; err != nil {
        return existingAuthor, err
    }
    existingAuthor = author
    existingAuthor.ID = id
    err := repo.DB.Save(&existingAuthor).Error
    return existingAuthor, err
}

func (repo *AuthorRepo) DeleteAuthor(id uint) error {
    return repo.DB.Delete(&models.Author{}, id).Error
}

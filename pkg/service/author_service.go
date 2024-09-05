package service

import (
	"author_book_api/pkg/models"
	"author_book_api/pkg/repository"
)

type AuthorService struct {
    repo *repository.AuthorRepo
}

func NewAuthorService(repo *repository.AuthorRepo) *AuthorService {
    return &AuthorService{repo: repo}
}

func (s *AuthorService) GetAllAuthors() ([]models.Author, error) {
    return s.repo.GetAllAuthors()
}

func (s *AuthorService) GetAuthorByID(id uint) (models.Author, error) {
    return s.repo.GetByID(id)
}

func (s *AuthorService) CreateAuthor(author models.Author) (models.Author, error) {
	return s.repo.CreateAuthor(author)
}

func (s *AuthorService) UpdateAuthor(id uint,author models.Author) (models.Author, error) {
	return s.repo.UpdateAuthor(id,author)
}

func (s *AuthorService) DeleteAuthor(id uint) error {
	return s.repo.DeleteAuthor(id)
}




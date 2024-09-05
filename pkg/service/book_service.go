package service

import (
	"author_book_api/pkg/models"
	"author_book_api/pkg/repository"
)

type BookService struct {
    repo *repository.BookRepo
}

func NewBookService(repo *repository.BookRepo) *BookService {
    return &BookService{repo: repo}
}

func (s *BookService) GetAllBooks() ([]models.Book, error) {
    return s.repo.GetAll()
}

func (s *BookService) GetBookByID(id uint) (models.Book, error) {
    return s.repo.GetByID(id)
}
func (s *BookService) CreateBook(book models.Book) (models.Book, error) {
	return s.repo.Create(book)
}

func (s *BookService) UpdateBook(id uint, book models.Book) (models.Book, error) {
	return s.repo.Update(id,book)
}

func (s *BookService) DeleteBook(id uint) error {
	return s.repo.Delete(id)
}


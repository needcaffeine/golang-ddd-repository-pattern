package service

import (
	"golang-ddd-repository-pattern/internal/repository"
	"golang-ddd-repository-pattern/model"
)

type BookstoreService struct {
	BookStoreRepository *repository.BookstoreRepository
}

func NewBookstoreService(bookStoreRepository *repository.BookstoreRepository) *BookstoreService {
	return &BookstoreService{
		BookStoreRepository: bookStoreRepository,
	}
}

func (s *BookstoreService) AddToInventory(title string, author string) error {
	book := model.NewBook(title, author)
	return s.BookStoreRepository.AddToInventory(book)
}

func (s *BookstoreService) OneById(id int) (model.Book, error) {
	return s.BookStoreRepository.OneById(id)
}

func (s *BookstoreService) OneByTitle(title string) (model.Book, error) {
	return s.BookStoreRepository.OneByTitle(title)
}

func (s *BookstoreService) AllBooks() ([]model.Book, error) {
	return s.BookStoreRepository.All()
}

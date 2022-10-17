package handler

import (
	"fmt"
	"golang-ddd-repository-pattern/internal/service"
	"log"
	"net/http"
)

type Handler struct {
	bookStoreService *service.BookstoreService
}

func NewHandler(bookStoreService *service.BookstoreService) *Handler {
	return &Handler{
		bookStoreService: bookStoreService,
	}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

// List all the books in our store. (This is a fictional store, don't judge.)
func (h *Handler) ListBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.bookStoreService.AllBooks()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Books in our store:\n")
	for _, book := range books {
		fmt.Fprintf(w, "%d: %s by %s\n", book.Id, book.Title, book.Author)
	}
}

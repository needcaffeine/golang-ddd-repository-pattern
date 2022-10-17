package main

import (
	"log"
	"net/http"

	"golang-ddd-repository-pattern/internal/database"
	"golang-ddd-repository-pattern/internal/handler"
	"golang-ddd-repository-pattern/internal/repository"
	"golang-ddd-repository-pattern/internal/service"

	_ "github.com/mattn/go-sqlite3"
)

var bookStoreService *service.BookstoreService

// This is a fictional bookstore management system.
func main() {
	// Get a database connection.
	db, err := database.InitDatabase("bookstore")
	if err != nil {
		log.Fatal(err)
	}

	// Dependency inject our database -> repository -> service -> controller
	bookStoreRepository := repository.NewBookstoreRepository(db)
	bookStoreService = service.NewBookstoreService(bookStoreRepository)

	// Populate our store with some books.
	bookStoreService.AddToInventory("The Hobbit", "J.R.R. Tolkien")
	bookStoreService.AddToInventory("The Fellowship of the Ring", "J.R.R. Tolkien")
	bookStoreService.AddToInventory("The Two Towers", "J.R.R. Tolkien")
	bookStoreService.AddToInventory("The Return of the King", "J.R.R. Tolkien")

	// Our books controller / handler will need access to the service
	booksHandler := handler.NewHandler(bookStoreService)

	// Start our server.
	http.HandleFunc("/", booksHandler.Home)
	http.HandleFunc("/books", booksHandler.ListBooks)

	http.ListenAndServe(":3000", nil)
}

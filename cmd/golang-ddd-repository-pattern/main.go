package main

import (
	"fmt"
	"log"

	"golang-ddd-repository-pattern/internal/database"
	"golang-ddd-repository-pattern/internal/repository"
	"golang-ddd-repository-pattern/internal/service"

	_ "github.com/mattn/go-sqlite3"
)

// This is a fictional bookstore management system.
func main() {
	// Get a database connection.
	db, err := database.InitDatabase("bookstore")
	if err != nil {
		log.Fatal(err)
	}

	// Dependency inject our database -> repository -> service -> controller
	bookStoreRepository := repository.NewBookstoreRepository(db)
	bookStoreService := service.NewBookstoreService(bookStoreRepository)

	// Populate our store with some books.
	bookStoreService.AddToInventory("The Hobbit", "J.R.R. Tolkien")
	bookStoreService.AddToInventory("The Fellowship of the Ring", "J.R.R. Tolkien")
	bookStoreService.AddToInventory("The Two Towers", "J.R.R. Tolkien")
	bookStoreService.AddToInventory("The Return of the King", "J.R.R. Tolkien")

	// List all the books in our store. (This is a fictional store, don't judge.)
	fmt.Println("Books in our store:")
	books, err := bookStoreService.AllBooks()
	if err != nil {
		log.Fatal(err)
	}
	for _, book := range books {
		log.Printf("%d: %s by %s", book.Id, book.Title, book.Author)
	}
}

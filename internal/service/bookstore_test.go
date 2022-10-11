package service

import (
	"golang-ddd-repository-pattern/internal/database"
	"golang-ddd-repository-pattern/internal/repository"
	"log"
	"os"
	"testing"

	"github.com/tj/assert"
)

var bookstoreService *BookstoreService

func TestMain(m *testing.M) {
	db, err := database.InitDatabase("../../bookstore_test")
	if err != nil {
		log.Fatal(err)
	}

	bookStoreRepository := repository.NewBookstoreRepository(db)
	bookstoreService = NewBookstoreService(bookStoreRepository)

	code := m.Run()
	os.Exit(code)
}

func TestInventoryOperations(t *testing.T) {
	suite := []struct {
		name   string
		title  string
		author string
	}{
		{
			name:   "AllRequiredParams",
			title:  "The Martian",
			author: "Andy Weir",
		},
		{
			name:   "AllRequiredParams",
			title:  "Artemis",
			author: "Andy Weir",
		},
		{
			name:   "AllRequiredParams",
			title:  "Hail Mary Project",
			author: "Andy Weir",
		},
	}

	for _, testCase := range suite {
		t.Run(testCase.name, func(t *testing.T) {
			err := bookstoreService.AddToInventory(testCase.title, testCase.author)
			assert.NoError(t, err)

			bookFromInventory, err := bookstoreService.OneByTitle(testCase.title)
			assert.NoError(t, err)
			assert.Equal(t, testCase.title, bookFromInventory.Title)
			assert.Equal(t, testCase.author, bookFromInventory.Author)

			bookFromInventory, err = bookstoreService.OneById(bookFromInventory.Id)
			assert.NoError(t, err)
			assert.Equal(t, testCase.title, bookFromInventory.Title)
			assert.Equal(t, testCase.author, bookFromInventory.Author)
		})
	}

	// Get all books from inventory and check that the number of books matches the number of test cases.
	booksFromInventory, err := bookstoreService.AllBooks()
	assert.NoError(t, err)
	assert.Len(t, booksFromInventory, len(suite))
}

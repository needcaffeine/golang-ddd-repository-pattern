package repository

import (
	"golang-ddd-repository-pattern/internal/database"
	"golang-ddd-repository-pattern/model"

	"log"
	"testing"

	"github.com/tj/assert"
)

func TestInventoryOperations(t *testing.T) {
	db, err := database.InitDatabase("../../bookstore_test")
	if err != nil {
		log.Fatal(err)
	}

	bookStoreRepository := NewBookstoreRepository(db)

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
			book := model.NewBook(testCase.title, testCase.author)

			// Add book to inventory.
			err := bookStoreRepository.AddToInventory(book)
			assert.NoError(t, err)

			// Get book from inventory by title and check that it matches.
			bookFromInventory, err := bookStoreRepository.OneByTitle(book.Title)
			assert.NoError(t, err)
			assert.Equal(t, testCase.title, bookFromInventory.Title)
			assert.Equal(t, testCase.author, bookFromInventory.Author)

			// Get book from inventory by id and check that it matches.
			bookFromInventory, err = bookStoreRepository.OneById(bookFromInventory.Id)
			assert.NoError(t, err)
			assert.Equal(t, testCase.title, bookFromInventory.Title)
			assert.Equal(t, testCase.author, bookFromInventory.Author)
		})
	}

	// Get all books from inventory.
	booksFromInventory, err := bookStoreRepository.All()
	assert.NoError(t, err)
	assert.Len(t, booksFromInventory, len(suite))
}

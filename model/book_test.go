package model

import (
	"testing"

	"github.com/tj/assert"
)

func TestNewBook(t *testing.T) {
	book := NewBook("The Martian", "Andy Weir")

	assert.Equal(t, book.Title, "The Martian")
	assert.Equal(t, book.Author, "Andy Weir")
}

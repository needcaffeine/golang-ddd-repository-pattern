package repository

import (
	"database/sql"

	"golang-ddd-repository-pattern/model"
)

type BookstoreRepository struct {
	db *sql.DB
}

func NewBookstoreRepository(db *sql.DB) *BookstoreRepository {
	return &BookstoreRepository{
		db: db,
	}
}

func (r *BookstoreRepository) AddToInventory(book *model.Book) error {
	_, err := r.db.Exec("INSERT INTO books (title, author) VALUES (?, ?)", book.Title, book.Author)
	return err
}

func (r *BookstoreRepository) OneById(id int) (model.Book, error) {
	var book model.Book
	err := r.db.QueryRow("SELECT id, title, author FROM books WHERE id = ?", id).Scan(&book.Id, &book.Title, &book.Author)

	return book, err
}

func (r *BookstoreRepository) OneByTitle(title string) (model.Book, error) {
	var book model.Book
	err := r.db.QueryRow("SELECT id, title, author FROM books WHERE title = ?", title).Scan(&book.Id, &book.Title, &book.Author)

	return book, err
}

func (r *BookstoreRepository) All() ([]model.Book, error) {
	rows, err := r.db.Query("SELECT id, title, author FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		var book model.Book
		if err := rows.Scan(&book.Id, &book.Title, &book.Author); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

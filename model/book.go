package model

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func NewBook(title string, author string) *Book {
	return &Book{
		Title:  title,
		Author: author,
	}
}

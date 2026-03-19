package books

import (
	"fmt"
	"maps"
	"slices"
)

type Book struct {
	ID     string
	Title  string
	Author string
	Copies int
}

var Catalog = map[string]Book{
	"abc": {
		Title:  "In the Company of Cheerful Ladies",
		Author: "Alexander McCall Smith",
		Copies: 1,
		ID:     "abc",
	},
	"xyz": {
		Title:  "White Heat",
		Author: "Dominic Sandbrook",
		Copies: 2,
		ID:     "xyz",
	},
}

func BookToString(book Book) string {
	return fmt.Sprintf("%v by %v (copies: %v)",
		book.Title, book.Author, book.Copies)
}

func GetAllBooks() []Book {
	return slices.Collect(maps.Values(Catalog))
}

func GetBook(id string) (Book, bool) {
	book, ok := Catalog[id]
	return book, ok
}

func AddBook(book Book) {
	Catalog[book.ID] = book
}

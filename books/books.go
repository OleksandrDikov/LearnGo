package books

import "fmt"

type Book struct {
	ID     string
	Title  string
	Author string
	Copies int
}

var catalog = []Book{
	{
		ID:     "abc",
		Title:  "In the Company of Cheerful Ladies",
		Author: "Alexander McCall Smith",
		Copies: 1,
	},
	{
		ID:     "xyz",
		Title:  "White Heat",
		Author: "Dominic Sandbrook",
		Copies: 2,
	},
}

func BookToString(book Book) string {
	return fmt.Sprintf("%v by %v (copies: %v)",
		book.Title, book.Author, book.Copies)
}

func GetAllBooks() []Book {
	return catalog
}

func GetBook(id string) (Book, bool) {
	for _, book := range GetAllBooks() {
		if book.ID == id {
			return book, true
		}
	}
	return Book{}, false
}

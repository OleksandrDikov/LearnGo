package books

import "fmt"

type Book struct {
	Title  string
	Author string
	ID string
	Copies int
}

var catalog = []Book{
	{
		Title:  "In the Company of Cheerful Ladies",
		Author: "Alexander McCall Smith",
		Copies: 1,
		ID: "abc",
	},
	{
		Title:  "White Heat",
		Author: "Dominic Sandbrook",
		Copies: 2,
		ID: "qwe",
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
	for _, book := range catalog {
		if book.ID == id {
			return book, true
		}
	}
	return Book{}, false
}
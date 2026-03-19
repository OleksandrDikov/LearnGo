package main

import (
	"books"
	"fmt"
	"maps"
)

func main() {
	catalog := books.GetCatalog()

	for _, book := range books.GetAllBooks(catalog) {
		fmt.Println(books.BookToString(book))
	}

	for book := range maps.Values(catalog) {
		fmt.Println(book)
		fmt.Println(books.BookToString(book))
	}
}

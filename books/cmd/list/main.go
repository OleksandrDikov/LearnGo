package main

import (
	"books"
	"fmt"
	"maps"
)

func main() {
	// book := books.Book{
	// 	Title:  "Engineering in Plain Sight",
	// 	Author: "Grady Hillhouse",
	// 	Copies: 2,
	// }
	// fmt.Println(books.BookToString(book))

	for _, book := range books.GetAllBooks() {
		fmt.Println(books.BookToString(book))
	}

	for book := range maps.Values(books.Catalog) {
		fmt.Println(book)
		fmt.Println(books.BookToString(book))
	}
}

package main

import (
	"books"
	"fmt"
)

func main() {
	fmt.Println(books.GetAllBooks())
	for _, v := range(books.GetAllBooks()) {
		fmt.Println(books.BookToString(v))
	}
}
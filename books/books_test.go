package books_test

import (
	"books"
	"cmp"
	"slices"
	"testing"
)

func getTestCatalog() map[string]books.Book {
	return map[string]books.Book{
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
}

func TestBookToString_FormatsBookInfoAsString(t *testing.T) {
	input := books.Book{
		Title:  "Sea Room",
		Author: "Adam Nicolson",
		Copies: 2,
	}
	want := "Sea Room by Adam Nicolson (copies: 2)"
	got := books.BookToString(input)
	if want != got {
		t.Fatalf("want %q, got %q", want, got)
	}
}

func TestGetAllBooks_ReturnsAllBooks(t *testing.T) {
	want := []books.Book{
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
	catalog := getTestCatalog()
	got := books.GetAllBooks(catalog)
	slices.SortFunc(got, func(a, b books.Book) int {
		return cmp.Compare(a.Author, b.Author)
	})
	if !slices.Equal(want, got) {
		t.Fatalf("want %#v, got %#v", want, got)
	}
}

func TestGetBook_FindBookInCatalogByID(t *testing.T) {
	t.Parallel()
	want := books.Book{
		ID:     "abc",
		Title:  "In the Company of Cheerful Ladies",
		Author: "Alexander McCall Smith",
		Copies: 1,
	}
	catalog := getTestCatalog()
	got, ok := books.GetBook(catalog, "abc")
	if !ok {
		t.Fatal("book not found")
	}
	if got != want {
		t.Fatalf("Want: %v, Got: %v", want, got)
	}
}

func TestGetBook_ReturnFalseWhenBookNotFound(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	_, ok := books.GetBook(catalog, "nonexisting ID")
	if ok {
		t.Fatal("want false for nonexistent ID, got true")
	}
}

func TestAddBook_AddBookToTheCatalog(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	_, ok := books.GetBook(catalog, "aab")
	if ok {
		t.Fatal("book already exist")
	}
	books.AddBook(catalog, books.Book{
		ID:     "aab",
		Title:  "Test Book",
		Author: "Alexander Smith",
		Copies: 1,
	})
	_, ok = books.GetBook(catalog, "aab")
	if !ok {
		t.Fatal("book not found")
	}
}

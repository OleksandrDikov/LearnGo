package books_test

import (
	"books"
	"slices"
	"testing"
)

func TestGetAllBooks_ReturnsAllBooks(t *testing.T) {
	t.Parallel()
	want := []books.Book{
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
	got := books.GetAllBooks()
	if !slices.Equal(want, got) {
		t.Fatalf("want %#v, got %#v", want, got)
	}
}

func TestGetBook__FindsBookInCatalogByID(t *testing.T) {
	t.Parallel()
	want := books.Book{
			Title:  "In the Company of Cheerful Ladies",
			Author: "Alexander McCall Smith",
			Copies: 1,
			ID: "abc",
	}
	got, ok := books.GetBook("abc")
	if !ok {
		t.Fatalf("book not found")
	}
	if want != got {
        t.Fatalf("want %#v, got %#v", want, got)
    }
}

func TestGetBook_ReturnsFalseWhenBookNotFound(t *testing.T) {
    t.Parallel()
    _, ok := books.GetBook("nonexistent ID")
    if ok {
        t.Fatal("want false for nonexistent ID, got true")
    }
}
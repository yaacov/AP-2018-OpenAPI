// The gen tool automatically generates routing code and openapi3 documents
// Gen can also generate client code
//
// Basic usage:
//   1. Install gen tool `go get -u -v github.com/wzshiming/gen/cmd/gen`
//   2. Add gen tool to $PATH
//   3. Execute it `gen run github.com/yaacov/AP-2018-OpenAPI/bookstore/v2/service`
//   4. Open http://127.0.0.1:8080/swagger/?url=./openapi.json# with your browser.

//go:generate gen route github.com/yaacov/AP-2018-OpenAPI/bookstore/v2/service
//go:generate gen openapi -o ../openapi/openapi.json github.com/yaacov/AP-2018-OpenAPI/bookstore/v2/service
//go:generate gen client -o ../client/client_gen.go github.com/yaacov/AP-2018-OpenAPI/bookstore/v2/service

package service

import (
	"fmt"
)

// Book in our book strore.
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Tag    string `json:"tag"`
	Pages  int    `json:"pages"`
}

var size int
var books = map[int]*Book{}

// GetBooks #route:"GET /books/"#
func GetBooks(limit int, filter string) (bookList []*Book, err error) {
	bookList = []*Book{}

	for i := 1; i <= size; i++ {
		bookList = append(bookList, books[i])
	}

	return
}

// CreateBook #route:"POST /books/"#
func CreateBook(book *Book) (newBook *Book, err error) {
	size++
	book.ID = size
	books[size] = book

	return book, nil
}

// GetBook #route:"GET /books/{bookID}"#
func GetBook(bookID int) (book *Book, err error) {
	if b, ok := books[bookID]; ok {
		return b, nil
	}

	return nil, fmt.Errorf("Book doesn't exist")
}

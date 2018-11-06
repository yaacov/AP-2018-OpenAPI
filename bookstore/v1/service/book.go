// The gen tool automatically generates routing code and openapi3 documents
// Gen can also generate client code
//
// Basic usage:
//   1. Install gen tool `go get -u -v github.com/wzshiming/gen/cmd/gen`
//   2. Add gen tool to $PATH
//   3. Execute it `gen run github.com/yaacov/AP-2018-OpenAPI/bookstore/v1/service`
//   4. Open http://127.0.0.1:8080/swagger/?url=./openapi.json# with your browser.

//go:generate gen route github.com/yaacov/AP-2018-OpenAPI/bookstore/v1/service
//go:generate gen openapi -o ../openapi/openapi.json github.com/yaacov/AP-2018-OpenAPI/bookstore/v1/service
//go:generate gen client -o ../client/client_gen.go github.com/yaacov/AP-2018-OpenAPI/bookstore/v1/service

package service

// Book in our book strore.
type Book struct {
	ID     int
	Title  string
	Author string
	Tag    string
	Pages  int
}

var size = 3
var books = []*Book{
	&Book{ID: 1, Title: "Best Book Ever", Pages: 4},
	&Book{ID: 2, Title: "Nice Book", Pages: 128},
	&Book{ID: 3, Title: "Even Beter Book", Pages: 225},
}

// GetBooks #route:"GET /books/"#
func GetBooks(limit int) (bookList []*Book, err error) {
	return books, nil
}

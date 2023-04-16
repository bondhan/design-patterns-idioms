package main

import "fmt"

type Book struct {
	Title     string
	Author    string
	Year      int
	Publisher string
}

func (b Book) String() string {
	return fmt.Sprintf("%s, written by %s in %d, published by %s", b.Title, b.Author, b.Year, b.Publisher)
}

type BookCollection struct {
	books []Book
}

func (bc *BookCollection) AddBook(book Book) {
	bc.books = append(bc.books, book)
}

func (bc *BookCollection) GetBooks() []Book {
	return bc.books
}

func main() {
	book1 := Book{Title: "The Go Programming Language", Author: "Alan A. A. Donovan and Brian W. Kernighan", Year: 2015, Publisher: "Addison-Wesley Professional"}
	book2 := Book{Title: "Effective Go", Author: "The Go Authors", Year: 2014, Publisher: "golang.org"}

	bookCollection := BookCollection{}
	bookCollection.AddBook(book1)
	bookCollection.AddBook(book2)

	books := bookCollection.GetBooks()

	for _, book := range books {
		fmt.Println(book)
	}
}

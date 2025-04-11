// challenges/types-composite/begin/main.go
package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

// struct types for author and book
type author struct{ name string }

type book struct {
	title string
	author
}

// define a library as a map from author name to list of books
type library map[string][]book

// define addBook to append a book to the library’s list
func (l library) addBook(b book) {
	authorName := b.author.name
	l[authorName] = append(l[authorName], b)
}

func (b book) display() string {
	return fmt.Sprintf("%s by %s", b.title, b.author.name)
}

// define a lookupByAuthor function to find books by an author's name
func (l library) lookupByAuthor(authorName string) []book {
	return l[authorName]
}

func main() {
	// create a new library
	lib := make(library) // or just: library{}

	// add 2 books to the library by the same author
	king := author{name: "Stephen King"}
	lib.addBook(book{author: king, title: "The Institute"})
	lib.addBook(book{author: king, title: "The Shining"})

	// add 1 book to the library by a different author
	lib.addBook(book{author: author{name: "Marie Curie"}, title: "Radium"})

	// dump the library with spew
	spew.Dump(lib)

	// lookup books by known author in the library
	searchTerm := king.name
	searchResults := lib.lookupByAuthor(searchTerm)
	fmt.Printf("Books by %s: %v\n", searchTerm, searchResults)

	if len(searchResults) > 0 {
		// print out the first book's title and its author's name
		fmt.Println("1st book matching '" + searchTerm + "': " + searchResults[0].display())
	} else {
		fmt.Println("No book matches '" + searchTerm + "'")
	}
	// if we look up an unknown author, we get an empty array back
	// spew.Dump(lib.lookupByAuthor("foo"))
	// ↑ that is why we added a length check on the search results above
}

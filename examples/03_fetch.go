package main

import (
	"bytes"
	"fmt"

	"github.com/toudi/yti"
)

func main() {
	type Book struct {
		Title  string `yaml:"title"`
		Author string `yaml:"author"`
	}
	// in this example we'll demonstrate how to use .Fetch() methods
	// to retrieve lists of items:
	data := bytes.NewBufferString(`
- title: The lord of the rings
  author: J.R.R. Tolkien
- title: Feynman's lectures of physics, Vol 1
  author: Richard P. Feynman
- title: Feynman's lectures of physics, Vol 2
  author: Richard P. Feynman
`)
	table, err := yti.Open[Book](data, &yti.TableOptions[Book]{
		// of course, you don't have to create any indices
		// however it's usually worth it to make the retrieve operations
		// faster.
		Indices: map[string]yti.Indexer[Book]{
			"author": func(item Book) interface{} {
				return item.Author
			},
		},
	})
	if err != nil {
		panic(err)
	}

	books_by_feynman := table.Fetch(func(item Book) bool {
		return item.Author == "Richard P. Feynman"
	})
	fmt.Printf("found books: %+v\n", books_by_feynman)
	// now, let's try to retrieve the same books but with index:
	books_by_feynman_with_index, err := table.FetchByIndexValue("author", "Richard P. Feynman")
	if err != nil {
		panic(err)
	}
	if len(books_by_feynman) != len(books_by_feynman_with_index) {
		fmt.Printf("the index malfunctioned\n")
	}
	// that's not exactly rocket science, is it ?
}

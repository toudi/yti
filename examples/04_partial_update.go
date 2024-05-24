package main

import (
	"bytes"
	"fmt"

	"github.com/toudi/yti"
)

func main() {
	// the previous examples used plain structs. However, if you
	// use pointers there are some interesting consequences of that.
	// again, let's represent the type for our data:

	type Book struct {
		Title  string `yaml:"title"`
		Read   bool   `yaml:"read"`
		Author string `yaml:"author"`
	}

	// and the data itself:

	data := bytes.NewBufferString(`
- title: Cooking recipies
  author: Remy the chef
- title: C++, the essentials
  author: The C++ comitee
- title: Geography of the North pole
  author: somebody
`)

	table, err := yti.Open[*Book](data, &yti.TableOptions[*Book]{
		// again, the indices are not required but I wanted to demonstrate
		// the nice property in that if you use the Update function
		// the changes are also reflected in the index.
		Indices: map[string]yti.Indexer[*Book]{
			"read": func(item *Book) interface{} {
				return item.Read
			},
		},
	})

	if err != nil {
		panic(err)
	}

	// make sure that initially none of the books are marked as read
	read_books := table.Fetch(func(item *Book) bool {
		return item.Read
	})
	if len(read_books) > 0 {
		panic("Something is wrong with the code.")
	}
	table.Update(func(item *Book) bool {
		if item.Title == "Cooking recipies" {
			item.Read = true
			// return true to indicate that we've
			// updated this row
			return true
		}
		// return false to indicate that no changes were made
		return false
	})
	// now we can use the index to retrieve all the books that we've written:
	read_books, err = table.FetchByIndexValue("read", true)
	if err != nil {
		panic("could not retrieve items by index")
	}
	if len(read_books) != 1 {
		panic("something is wrong with the code")
	}
	fmt.Printf("Here are all the books that we've read:\n")
	for _, book := range read_books {
		fmt.Printf("-> %+v\n", book)
	}
}

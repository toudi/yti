package main

import (
	"bytes"
	"fmt"

	"github.com/toudi/yti"
)

func main() {
	// as you saw from the previous example, the `Get` method wasn't exactly
	// efficient. However, we can help the library by using indices. Here's
	// how to do that:
	//
	// let's define our structure like we previously did:
	type Movie struct {
		Id   uint16 `yaml:"id"`
		Name string `yaml:"name"`
		Year uint16 `yaml:"year"`
	}
	data := bytes.NewBufferString(`
- id: 1
  name: The Shawshank Redemption
  year: 1994
- id: 2
  name: The Godfather
  year: 1972
- id: 3
  name: The Dark Knight
  year: 2008
`)

	// now, let's try to get `The Shawshank Redemption`:
	table, err := yti.Open[Movie](data, &yti.TableOptions[Movie]{
		Indices: map[string]yti.Indexer[Movie]{
			"id": func(item Movie) interface{} {
				return item.Id
			},
		},
	})
	if err != nil {
		panic(err)
	}
	// now let's try the .GetByIndex call:
	movie, _ := table.GetByIndex("id", 1)
	fmt.Printf("Found %s\n", movie.Name)

	// now let's try to make the program fault. We will do so by inserting some
	// random movie with an existing id:
	table.Insert(Movie{Id: 1, Name: "some test movie", Year: 1962})

	// now we can make sure that both .Get and .GetByIndex will fail - though
	// .GetByIndex will fail faster:
	_, err = table.Get(func(item Movie) bool {
		return item.Id == 1
	})
	fmt.Printf(".Get() got error: %v\n", err)
	_, err = table.GetByIndex("id", 1)
	fmt.Printf(".GetByIndex() got error: %v\n", err)
}

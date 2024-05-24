package main

import (
	"bytes"
	"fmt"

	"github.com/toudi/yti"
)

func main() {
	// in this example, we're going to demonstrate two functions - Get and First.
	// However, first, let's look at the example data. These are
	// some movies that are taken from imdb's top 250 website
	// for demonstrational purposes.
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

	// we need to model the table into a struct of ours:

	type Movie struct {
		Id   uint16 `yaml:"id"`
		Name string `yaml:"name"`
		Year uint16 `yaml:"year"`
	}

	// now, let's try to get `The Shawshank Redemption`:
	table, err := yti.Open[Movie](data, nil)
	if err != nil {
		panic(err)
	}
	// the second parameter are the the table options which we will chose to ignore
	// at this moment

	fmt.Printf("table.Get() example\n")
	// let's proceed:
	movie, _ := table.Get(func(item Movie) bool {
		fmt.Printf("inspecting %s\n", item.Name)
		return item.Id == 1
	})

	fmt.Printf("Found %s\n", movie.Name)
	// Great, we've found our movie. But did you also notice that the program printed
	// the other movies despite the fact that the first one was already a match we were
	// looking for ?
	// That is because, Get checks if the criteria you've passed result in a unique
	// assignment. If you want to skip this check, use `First` instead:
	fmt.Printf("table.First() example\n")
	movie, _ = table.First(func(item Movie) bool {
		fmt.Printf("inspecting %s\n", item.Name)
		return item.Id == 1
	})
	fmt.Printf("Found %s\n", movie.Name)
}

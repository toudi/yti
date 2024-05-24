package main

import (
	"fmt"
	"os"

	"github.com/toudi/yti"
)

func main() {
	// all the previous examples used in-memory buffers but chances are
	// you'd want to work with files.

	type Book struct {
		Title string `yaml:"title"`
	}

	const filename = "some-file.yaml"
	defer os.Remove(filename)

	// by default, yti does not panic when opening the nonexisting file:
	_, err := yti.OpenFile[Book](filename, nil)
	if err != nil {
		panic("unexpected error")
	}

	// the reason behind this is that you may want to have a semi-existing file
	// (let's say - that exists in /tmp)
	// and then if it doesn't exist you'd just repopulate it with some known values

	// however you can indicate that the file must exist:
	_, err = yti.OpenFile[Book](filename, &yti.TableOptions[Book]{
		MustExist: true,
	})
	fmt.Printf("Open file returned an error: %v\n", err)

	// another reason for opening a file that may not exist is that yti will create one
	// for you when calling .Close()

	// if you use a directory name, and the said directory does not exist,
	// yti can also create it for you during .Close(), but only with MkDirs
	// option:
	_, err = yti.OpenFile[Book]("nested/directories/file.yaml", &yti.TableOptions[Book]{
		MustExist: true,
		MkDirs:    true,
	})
	fmt.Printf("Open file returned an error: %v\n", err)
}

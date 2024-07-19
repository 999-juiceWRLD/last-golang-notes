package main

import (
	"fmt"
	"path"
)

func _() {

	var someString string = "/bin/some-notes/note1.pdf"
	dir, str := path.Split(someString)
	fmt.Printf("DIR: %s | STR: %s\n", dir, str)
}
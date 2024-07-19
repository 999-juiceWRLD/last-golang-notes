package main

import (
	"fmt"
	"unicode/utf8"
)

func _() {
	name := "kahraman"
	fmt.Println(utf8.RuneCountInString(name))
}
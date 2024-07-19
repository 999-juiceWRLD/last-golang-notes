package main

import (
	"fmt"
	"os"
	"strings"
)

func _() {
	// var text string = os.Args[1]
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%d, %s and \n", i, strings.ToUpper(os.Args[i]))
	}
	// fmt.Println(strings.ToUpper(text))
}
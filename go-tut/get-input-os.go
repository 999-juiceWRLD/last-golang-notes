package main

import (
	"os"
	"fmt"
)

func _() {
	
	// for i := 0; i < len(os.Args); i++ {
	// 	fmt.Printf("%v ", os.Args[i])
	// }

	fmt.Printf("%#v\n", os.Args)

	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("%d --> %s \n", i, os.Args[i])
	}

	// RESULT: 
	// go build get-input-os.go
	// ./get-input-os this is so good!
	// []string{"./get-input-os", "this", "is", "so", "good!"}
	// 0 --> ./get-input-os 
	// 1 --> this 
	// 2 --> is 
	// 3 --> so 
	// 4 --> good!

}

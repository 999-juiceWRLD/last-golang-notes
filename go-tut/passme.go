package main

import (
	"fmt"
	"os"
	"os/user"
)

func _() {
	var username string = os.Args[1]
	userName, err := user.Current()
	if err != nil {
		fmt.Println("error 1")
	} else if userName.Name != username {
		fmt.Println("not valid")
	} else {
		fmt.Println("it's done")
	}
	
}

package main

import "fmt"

func _() {

	fmt.Println("naber");
	kahraman := User{"kahraman", "kcanik@gmail.com", false, 19};
	fmt.Printf("%+v\n", kahraman);
}

type User struct {
	Name	string
	Email	string
	Status 	bool
	Age		int
}

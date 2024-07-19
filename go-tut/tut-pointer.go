package main

import "fmt"

func printTen(num *int) {
	fmt.Println(num)
}

func returnTen(num int) *int {
	return &num;
}

func _() {
	var ten int = 10
	var tenPtr *int = &ten;
	printTen(&ten)
	printTen(tenPtr)
}
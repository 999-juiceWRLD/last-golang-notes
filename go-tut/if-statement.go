package main

import "fmt"

func getLength(email string) int {
	return len(email);
}

func _() {
	if length := getLength("kcaniksi"); length < 1 {
		fmt.Println("no, not at all");
	}
}

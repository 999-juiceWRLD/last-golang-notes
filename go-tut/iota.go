package main

import (
	"fmt"
)

// iota is a built-in constant generator which generates ever
// increasing numbers

func _() {
	const (
		monday = iota
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)

	fmt.Println(monday, tuesday, wednesday, thursday, 
				friday, saturday, sunday)
}

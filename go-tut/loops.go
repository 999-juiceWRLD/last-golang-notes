package main

import "fmt"

func _() {
	var i float32 = 0;
	var count int = 1;
	for ; i < 10; i += 0.5  {
		if count % 2 == 0 {
			count++
			continue
		} else {
			fmt.Printf("run %d times\n", count)
		}
		count++
	}
}

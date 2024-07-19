package main

import (
	"fmt"
	"time"
)

func _() {
	first_ch := make(chan string)
	second_ch := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		first_ch <- "sent to first_ch"
	}()

	go func() {
		time.Sleep(999 * time.Millisecond)
		second_ch <- "sent to second_ch"
	}()

	select {
	case receive_1 := <- first_ch:
		fmt.Printf("message received: %s\n", receive_1)
	case receive_2 := <- second_ch:
		fmt.Printf("message received: %s\n", receive_2)
	}
}
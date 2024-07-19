package main

import (
	"fmt"
	"time"
)

func sender(ch chan <- int) {
	for i := 0; i < 5; i++ {
		fmt.Println("Sending: ", i)
		ch <- i
		time.Sleep(time.Second)
	}
	close(ch)
}

func receiver(ch <- chan int) {
	for {
		val, ok := <- ch
		if !ok {
			fmt.Println("Channel closed, receiver exitting")
			return
		}
		fmt.Println("Receiving: ", val)
	}
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacci2(c2, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c2 <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func _() {
	dataChannel := make(chan int)
	go sender(dataChannel)
	go receiver(dataChannel)

	time.Sleep(6 * time.Second)
	fmt.Println("Main goroutine exiting.")

	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}

	fmt.Println("new one")

	c2 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c2)
		}
		quit <- 0
	}()
	fibonacci2(c2, quit)
}

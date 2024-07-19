package main

import (
	"fmt"
	"time"
	"sync"
)

// the *select* statement lets a goroutine wait on multiple
// communication operations. A *select* blocks until one of
// it's cases can run, then executes that case. Chooses random
// if multiple cases are ready.

var wg sync.WaitGroup

func _() {
	started := time.Now()
	foods := []string{"fettucini alfredo", "rice", "roast beef", "kobe steak"}
	wg.Add(len(foods))
	for _, el := range foods {
		go func(f string) {
			cook(&f)
			wg.Done()
		}(el) 
	}
	wg.Wait()
	fmt.Printf("done in %v\n", time.Since(started))
}

func cook(food *string) {
	fmt.Printf("cooking %s...\n", *food)
	time.Sleep(2 * time.Second)
	fmt.Printf("done cooking %s\n", *food)
}

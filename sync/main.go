package main

import (
	"fmt"
	"sync"
)

func main() {
	waitgroup()
}

func waitgroup() {
	// Initialize
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		// Add 1 to the counter
		wg.Add(1)
		go hello(i, &wg)
	}
	// Wait blocks till the value of the counter becomes 0
	wg.Wait()
}
func hello(id int, wg *sync.WaitGroup) {
	// Done decrements counter by one
	defer wg.Done()
	fmt.Printf("Hello from No %d!\n", id)
}

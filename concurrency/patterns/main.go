package main

import (
	"fmt"
)

func main() {
	// generator()
	multiplexer()
}

// Generator pattern
// Calls a function that returns a read-only channel
func generator() {
	ann := boring("Ann")
	joe := boring("Joe")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %s\n", <-ann)
		fmt.Printf("You say %s\n", <-joe) // This will have to wait for the above to be ready
	}
	fmt.Println("You're boring. I'm leaving")
}

func multiplexer() {
	m := fanIn(boring("Ann"), boring("Joe"))
	for i := 0; i < 10; i++ {
		fmt.Printf("You say: %s\n", <-m)
	}
}

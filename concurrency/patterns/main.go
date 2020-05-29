package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Uncomment the line below to run the generator pattern
	// generator()
	// Uncomment the line below to run the multiplexer pattern
	// multiplexer()
	// Uncomment the line below for sequencing restored
	// restoreSequencing()
	// multiplexerWithSelect()
	// multiplexerSelectWTimer()
	// selectWTotalTimeout(boring("joe"))
	// generatorWQuit()
	// generatorWCleanup()
	search()

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

func restoreSequencing() {
	msg := fanInRestored(boringRestored("Ann"), boringRestored("Joe"))
	for i := 0; i < 5; i++ {
		msg1 := <-msg
		fmt.Println(msg1.str)
		msg2 := <-msg
		fmt.Println(msg2.str)
		// I don't quite understand how this works
		// Since we are sharing the same bool channel between these
		// goroutines why does it matter where the 'true' is sent?
		msg1.wait <- true
		msg2.wait <- true
	}
}

func multiplexerWithSelect() {
	ann, joe := boring("Ann"), boring("Joe")
	c := fanInWithSelect(ann, joe)
	for i := 0; i < 10; i++ {
		fmt.Printf("You say: %s\n", <-c)
	}
}

func multiplexerSelectWTimer() {
	ann := boring("Ann")
	c := selectWithTimer(ann)
	for i := 0; i < 10; i++ {
		fmt.Printf("You say: %s\n", <-c)
	}
}

func generatorWQuit() {
	rand.Seed(time.Now().UnixNano())
	quit := make(chan bool)
	joe := boringWQuit("joe", quit)
	for i := rand.Intn(20); i >= 0; i-- {
		fmt.Println(<-joe)
	}
	quit <- true
}

func generatorWCleanup() {
	rand.Seed(time.Now().UnixNano())
	quit := make(chan string)
	joe := quitWCleanup("joe", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-joe)
	}
	quit <- "Quit!"
	// This will block and allow for the goroutine to perform any cleanup operations
	fmt.Printf("Joe says %s\n", <-quit)
}

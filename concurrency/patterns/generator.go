package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Generator function, returns a read-only channel
func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("Hello %s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

// Stops sending messages when quit signal received
func boringWQuit(s string, quit <-chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("Hello %s %d\n", s, i):
				// do nothing
			case <-quit:
				fmt.Println("Time over!")
				return
			}
		}
	}()
	return c
}

func quitWCleanup(s string, quit chan string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("Hello %s %d\n", s, i):
				// do nothing
			case <-quit:
				// cleanup any used resources
				cleanup()
				// signal that cleanup is finished
				quit <- "Cleanup finished"
				return
			}
		}
	}()
	return c
}

func cleanup() {
	fmt.Println("Cleaning up used resources..")
}

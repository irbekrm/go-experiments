package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str string
	// signaller field. Goroutines will block waiting on this
	// and thus will synchronize
	wait chan bool
}

// generator function
// returns a channel to which a goroutine will be sending messages
func boringRestored(msg string) <-chan Message {
	waitForIt := make(chan bool)
	c := make(chan Message)
	go func() {
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s: %d", msg, i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
			// This is now blocking
			<-waitForIt
		}
	}()
	return c
}

// multiplexer function
// takes a bunch of channels to which goroutines are sending messages
// returns a single channel that receives all the messages
func fanInRestored(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

package main

import (
	"fmt"
	"time"
)

func fanInWithSelect(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

func selectWithTimer(input <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input:
				c <- s
			// Timer for this round in the loop
			case <-time.After(time.Duration(500) * time.Millisecond):
				c <- fmt.Sprint("Too slow!")
			}
		}
	}()
	return c
}

func selectWTotalTimeout(input <-chan string) {
	// Timer for the total time the loop will run
	t := time.After(time.Second * 2)
	for {
		select {
		case m := <-input:
			fmt.Println(m)
		case <-t:
			fmt.Println("Time over!")
			return
		}
	}
}

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
			case <-time.After(time.Duration(500) * time.Millisecond):
				c <- fmt.Sprint("Too slow!")
			}
		}
	}()
	return c
}

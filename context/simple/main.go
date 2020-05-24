package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	// cancelOnStdin()
	cancelWithTimeout()
}

func cancelOnStdin() {
	// context.Background creates empty (root) context
	ctx := context.Background()
	// create a child context of the one above
	ctx, cancel := context.WithCancel(ctx)
	// the function called by the goroutine will call cancel,
	// if it receives anything on stdin
	go func() {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		cancel()
	}()
	sayItLater(ctx, time.Second*5, "Labdien")
}

func cancelWithTimeout() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	// important to still call cancel() to release resources (timers) allocated by context
	defer cancel()
	sayItLater(ctx, time.Second*6, "Hello")
}

func sayItLater(ctx context.Context, t time.Duration, s string) {
	select {
	case <-time.After(t):
		fmt.Println(s)
	// if the context has been cancelled or deadline exceeded, print the corresponding error
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

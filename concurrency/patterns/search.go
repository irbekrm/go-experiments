package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web = fakeSearch("Web")
	Image = fakeSearch("Image")
	Video = fakeSearch("Video")
)

type Result string

type Search func(string) Result

func Google(query string) (results []Result) {
	c := make(chan Result)
	// Launch all three queries in separate goroutines
	// so that they sleep in parallel and the total time is the longest sleep time
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}
	return
}

func search() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google("golang")
	since := time.Since(start)
	fmt.Println(results)
	fmt.Println(since)
}

func fakeSearch(kind string) Search {
	return func(query string) Result {
		n := time.Duration(rand.Intn(20))
		fmt.Printf("Sleeping for %d seconds in %s query\n", n, kind)
		time.Sleep(time.Second * n)
		return Result(fmt.Sprintf("The result for %s query about %s\n", kind, query))
	}
}


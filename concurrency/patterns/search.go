package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web    = fakeSearch("Web")
	Image  = fakeSearch("Image")
	Video  = fakeSearch("Video")
	Web1   = fakeSearch("Web")
	Image1 = fakeSearch("Image")
	Video1 = fakeSearch("Video")
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

func GoogleWTimeout(query string) (results []Result) {
	c := make(chan Result)
	t := time.After(time.Second * 10)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()
	for i := 0; i < 3; i++ {
		select {
		case <-t:
			fmt.Println("Timeout!")
			return
		case result := <-c:
			results = append(results, result)
		}
	}
	return
}

// Will search with replication(First accepts a bunch of servers) and timeout for the whole search
func GoogleWTimeoutAndReplication(query string) (results []Result) {
	c := make(chan Result)
	t := time.After(time.Second * 15)
	go func() { c <- First(query, Web, Web1) }()
	go func() { c <- First(query, Video, Video1) }()
	go func() { c <- First(query, Image, Image1) }()
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-t:
			fmt.Println("Timeout!")
			return
		}
	}
	return
}

func search() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	// results := Google("golang")
	// results := GoogleWTimeout("golang")
	results := GoogleWTimeoutAndReplication("golang")
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

// Takes a query function and a bunch of replica servers
// Queries the replica servers and returns the first result
// This prevents waiting around after having hit a particular slow server
func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

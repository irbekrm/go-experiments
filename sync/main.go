package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

func main() {
	Waitgroup()
	PrintWithPool([]byte(`{"name":"Aina"}`), os.Stdout)
	PrintWithoutPool([]byte(`{"name":"Elga"}`), os.Stdout)
}

func Waitgroup() {
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

type user struct {
	Name string
}

var uPool = sync.Pool{
	New: func() interface{} { return new(user) },
}

// PrintWithPool can be called multiple times in parallel
func PrintWithPool(data []byte, out io.Writer) {
	// instead of allocating a new user every time the function is called, take one from the pool
	u := uPool.Get().(*user)
	// return user to the pool
	defer uPool.Put(u)
	// set name to empty string
	u.Name = ""
	if err := json.Unmarshal(data, u); err != nil {
		log.Fatalf("could not unmarshal json: %v", err)
	}
	fmt.Fprintf(out, "User name is %s\n", u.Name)
}

// PrintWithoutPool can be called multiple times in parallel
func PrintWithoutPool(data []byte, out io.Writer) {
	u := user{}
	if err := json.Unmarshal(data, &u); err != nil {
		log.Fatalf("could not unmarshal json: %v", err)
	}
	fmt.Fprintf(out, "User name is %s\n", u.Name)
}

func hello(id int, wg *sync.WaitGroup) {
	// Done decrements counter by one
	defer wg.Done()
	fmt.Printf("Hello from No %d!\n", id)
}

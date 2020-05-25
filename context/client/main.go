package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/context"
)

func main() {
	// Create a new, empty root context
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	// Wrapper for NewRequestWithContext (also uses the background context)
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	// New request with the given context
	req = req.WithContext(ctx)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)
	}
	io.Copy(os.Stdout, res.Body)
}

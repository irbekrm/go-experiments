package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("handler started")
	log.Println("handler ended")
	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "hello")
	case <-ctx.Done():
		err := ctx.Err()
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

const SockAddr = "/tmp/echo.sock"

func reader(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf[:])
	if err != nil {
		return
	}
	fmt.Printf("Client got: %v\n", string(buf[0:n]))
}

func main() {
	c, err := net.Dial("unix", SockAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	go reader(c)
	if _, err = c.Write([]byte("hi")); err != nil {
		log.Fatalf("Write error: %v\n", err)
	}
	reader(c)
	time.Sleep(100 * time.Millisecond)
}

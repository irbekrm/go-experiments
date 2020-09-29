package main

import (
	"bytes"
	"encoding/gob"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/irbekrm/learning-go/grpc/tutorialjff/todo"
	// v2 version of Go protobuf, v1 is at github.com/golang/protobuf/proto
	"google.golang.org/protobuf/proto"
)

// Code examples from https://www.youtube.com/watch?v=_jQ3i_fyqGA&t=840s&ab_channel=justforfunc%3AProgramminginGo

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "missing subcommand: list or add")
		os.Exit(1)
	}
	var err error
	switch cmd := flag.Arg(0); cmd {
	case "list":
		err = list()
	case "add":
		err = add(strings.Join(flag.Args()[1:], " "))
	default:
		err = fmt.Errorf("unknown subcommand: %s", cmd)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

const dbPath = "mydb.pb"

// Writes proto encoded messages to file
func add(text string) error {
	task := &todo.Task{
		Text: text,
		Done: false,
	}
	// Encodes the task to wire format
	b, err := proto.Marshal(task)
	if err != nil {
		return fmt.Errorf("could not encode task: %v", err)
	}
	f, err := os.OpenFile(dbPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("could not open %s: %v", dbPath, err)
	}
	// Need to add the size of the message otherwise when decoding it (with protobuf.Unmarshal)
	// there will be no way of reading messages one by one
	// Can view the contents of the file with cat [FILE] | protoc --decode_raw]
	if err := gob.NewEncoder(f).Encode(int64(len(b))); err != nil {
		return fmt.Errorf("could not encode lenght of message")
	}
	if _, err := f.Write(b); err != nil {
		return fmt.Errorf("could not write task to file: %v", err)
	}
	if err := f.Close(); err != nil {
		return fmt.Errorf("could not close file %s: %v", dbPath, err)
	}

	return nil
}

// Reads proto-encoded messages from a file
func list() error {
	// No need to close the file if it was read with ioutil.ReadFile
	b, err := ioutil.ReadFile(dbPath)
	if err != nil {
		return fmt.Errorf("could not read %s: %v", dbPath, err)
	}
	for {
		if len(b) == 0 {
			return nil
		} else if len(b) < 4 {
			return fmt.Errorf("remaining odd %d bytes, what to do?", len(b))
		}
		// Problem- how to know when to stop? (i.e how to separate the messages)
		// To distinguish between different messages, could store the size of each
		// so that when reading you know that next X bytes correspond to one message?
		// Solution- read the stored size X of the next message to know that that next X bytes correspond to one message
		var length int64
		if err := gob.NewDecoder(bytes.NewReader(b[:4])).Decode(&length); err != nil {
			fmt.Errorf("could not decode message lenght: %v", err)
		}
		// Move past size
		b = b[4:]

		var task todo.Task
		// Parses the wire-format messages in b and places in task
		if err := proto.Unmarshal(b[:length], &task); err != nil {
			return fmt.Errorf("could not read task: %v", err)
		}
		// Move past the message
		b = b[length:]

		if task.Done {
			fmt.Println("done")
		} else {
			fmt.Println("not done")
		}
		fmt.Println(task.Text)
	}
	return nil
}

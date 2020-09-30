package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	"github.com/irbekrm/learning-go/grpc/tutorialjff/todo"
	// v2 version of Go protobuf, v1 is at github.com/golang/protobuf/proto

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

// Code examples from https://www.youtube.com/watch?v=_jQ3i_fyqGA&t=840s&ab_channel=justforfunc%3AProgramminginGo

type taskServer struct {
}

func main() {
	srv := grpc.NewServer()
	tasks := taskServer{}
	todo.RegisterTasksServer(srv, tasks)
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("could not listen to :8888: %v", err)
	}
	log.Fatal(srv.Serve(l))
}

type length int64

const (
	sizeOfLength = 8
	dbPath       = "mydb.pb"
)

var endianness = binary.LittleEndian

func (s taskServer) Add(ctx context.Context, text *todo.Text) (*todo.Task, error) {
	task := &todo.Task{
		Text: text.Text,
		Done: false,
	}
	// Encodes the task to wire format
	b, err := proto.Marshal(task)
	if err != nil {
		return nil, fmt.Errorf("could not encode task: %v", err)
	}
	f, err := os.OpenFile(dbPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("could not open %s: %v", dbPath, err)
	}
	if err := binary.Write(f, endianness, length(len(b))); err != nil {
		return nil, fmt.Errorf("could not encode lenght of message")
	}
	if _, err := f.Write(b); err != nil {
		return nil, fmt.Errorf("could not write task to file: %v", err)
	}
	if err := f.Close(); err != nil {
		return nil, fmt.Errorf("could not close file %s: %v", dbPath, err)
	}

	return task, nil
}

func (s taskServer) List(ctx context.Context, void *todo.Void) (*todo.TaskList, error) {
	// Reads proto-encoded messages from a file
	// No need to close the file if it was read with ioutil.ReadFile
	b, err := ioutil.ReadFile(dbPath)
	if err != nil {
		return nil, fmt.Errorf("could not read %s: %v", dbPath, err)
	}
	var tasks todo.TaskList
	for {
		if len(b) == 0 {
			return &tasks, nil
		} else if len(b) < sizeOfLength {
			return nil, fmt.Errorf("remaining odd %d bytes, what to do?", len(b))
		}
		var l length

		if err := binary.Read(bytes.NewReader(b[:sizeOfLength]), endianness, &l); err != nil {
			return nil, fmt.Errorf("could not decode message length: %v", err)
		}
		// Move past size
		b = b[sizeOfLength:]

		var task todo.Task
		// Parses the wire-format messages in b and places in task
		if err := proto.Unmarshal(b[:l], &task); err != nil {
			return nil, fmt.Errorf("could not read task: %v", err)
		}
		// Move past the message
		b = b[l:]
		tasks.Tasks = append(tasks.Tasks, &task)
	}
}

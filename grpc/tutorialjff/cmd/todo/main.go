package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/irbekrm/learning-go/grpc/tutorialjff/todo"
	"google.golang.org/grpc"
	// v2 version of Go protobuf, v1 is at github.com/golang/protobuf/proto
)

// Code examples from https://www.youtube.com/watch?v=_jQ3i_fyqGA&t=840s&ab_channel=justforfunc%3AProgramminginGo

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "missing subcommand: list or add")
		os.Exit(1)
	}
	var err error
	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to backend: %v", err)
	}
	client := todo.NewTasksClient(conn)
	ctx := context.Background()
	switch cmd := flag.Arg(0); cmd {
	case "list":
		err = list(ctx, client)
	case "add":
		err = add(ctx, client, strings.Join(flag.Args()[1:], " "))
	default:
		err = fmt.Errorf("unknown subcommand: %s", cmd)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func add(ctx context.Context, client todo.TasksClient, text string) error {
	t := todo.Text{Text: text}
	_, err := client.Add(ctx, &t)
	if err != nil {
		return fmt.Errorf("could not add task: %v", err)
	}
	return nil
}

func list(ctx context.Context, client todo.TasksClient) error {
	l, err := client.List(ctx, &todo.Void{})
	if err != nil {
		return fmt.Errorf("could not fetch tasks: %v", err)
	}
	for _, t := range l.Tasks {
		if t.Done {
			fmt.Println("done")
		} else {
			fmt.Println("not done")
		}
		fmt.Printf(" %s\n", t.Text)
	}
	return nil
}

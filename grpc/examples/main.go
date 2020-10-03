package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/irbekrm/learning-go/grpc/examples/language"
	"google.golang.org/protobuf/proto"
)

type length int64

const dbPath = "phrases.pb"

func main() {
	var (
		addCommand = flag.NewFlagSet("add", flag.ExitOnError)
		langPtr    = addCommand.String("lang", "", "language")
		textPtr    = addCommand.String("text", "", "text")
	)

	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "missing subcommand: add")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "add":
		addCommand.Parse(os.Args[2:])
		if err := add(*textPtr, *langPtr); err != nil {
			fmt.Fprintf(os.Stderr, "could not add text %s: %v", *textPtr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s", os.Args[1])
		os.Exit(1)
	}
}

func add(text, lang string) error {
	val, ok := language.Phrase_Language_value[lang]
	if !ok {
		return fmt.Errorf("language %s not supported", lang)
	}
	l := language.Phrase_Language(val)
	phrase := language.Phrase{
		Text:     text,
		Language: l,
	}
	// wire-encode the phrase
	b, err := proto.Marshal(&phrase)

	f, err := os.OpenFile(dbPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("could not open %s: %v", dbPath, err)
	}
	_, err = f.Write(b)
	if err != nil {
		return fmt.Errorf("could not write task to file: %v", err)
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("could not close file %s: %v", dbPath, err)
	}
	fmt.Printf("Phrase is %#v", phrase)
	return nil
}

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/irbekrm/learning-go/grpc/examples/language"
	"google.golang.org/protobuf/proto"
)

type length int64

const (
	phrasesDb = "phrases.db"
	litDb     = "lit.db"
)

func main() {
	var (
		addPhraseCommand = flag.NewFlagSet("add_phrase", flag.ExitOnError)
		langPtr          = addPhraseCommand.String("lang", "", "language")
		textPtr          = addPhraseCommand.String("text", "", "text")
		addLitCommand    = flag.NewFlagSet("add_lit", flag.ExitOnError)
		kindPtr          = addLitCommand.String("kind", "", "kind: novel or poem")
		namePtr          = addLitCommand.String("name", "", "name of the piece")
		metrePtr         = addLitCommand.String("metre", "", "metre for kind poem")
		charactersPtr    = addLitCommand.String("chars", "", "comma separated characters for kind novel")
	)

	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "missing subcommand: add_phrase")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "add_phrase":
		addPhraseCommand.Parse(os.Args[2:])
		if err := addPhrase(*textPtr, *langPtr); err != nil {
			fmt.Fprintf(os.Stderr, "could not add phrase %s: %v", *textPtr, err)
			os.Exit(1)
		}
	case "add_lit":
		addLitCommand.Parse(os.Args[2:])
		if err := addLiteraryPiece(*namePtr, *kindPtr, *metrePtr, *charactersPtr); err != nil {
			fmt.Fprintf(os.Stderr, "could not add literary piece %s: %v", *namePtr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s", os.Args[1])
		os.Exit(1)
	}
}

func addPhrase(text, lang string) error {
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
	if err != nil {
		return fmt.Errorf("could not encode %v: %v", phrase, err)
	}
	return writeFile(b, phrasesDb)
}

func addLiteraryPiece(name, kind, metre, characters string) error {
	var lp language.LiteraryPiece
	switch kind {
	case "poem":
		p := language.LiteraryPiece_Poem{Metre: metre}
		pp := language.LiteraryPiece_Poem_{Poem: &p}
		lp = language.LiteraryPiece{Name: name, Type: &pp}
	case "novel":
		n := language.LiteraryPiece_Novel{Characters: characters}
		nn := language.LiteraryPiece_Novel_{Novel: &n}
		lp = language.LiteraryPiece{Name: name, Type: &nn}
	default:
		return fmt.Errorf("unknown type: %s", kind)
	}
	// wire-encode the struct
	b, err := proto.Marshal(&lp)
	if err != nil {
		return fmt.Errorf("could not encode %v: %v", lp, err)
	}
	return writeFile(b, litDb)
}

func writeFile(data []byte, filename string) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("could not open %s: %v", filename, err)
	}
	_, err = f.Write(data)
	if err != nil {
		return fmt.Errorf("could not write task to file: %v", err)
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("could not close file %s: %v", filename, err)
	}
	return nil
}

package main

import (
	"go/parser"
	"go/token"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	log.SetFlags(0)

	if len(os.Args) != 2 {
		log.Fatalf("expected path to file")
	}

	fs := token.NewFileSet()
	file, err := parser.ParseFile(fs, os.Args[1], nil, parser.AllErrors)
	if err != nil {
		log.Fatalf("failed to parse %s: %v", os.Args[1], err)
	}

	spew.Dump(file)
}

package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/pieroshka/gf/internal/analysis"
	"github.com/pieroshka/gf/internal/compile"
)

func main() {
	inPath := flag.String("in", "", "path to the Brainfuck source file (use '-' for stdin)")
	outPath := flag.String("out", "out/main.go", "path to the output file")
	flag.Parse()

	if *inPath == "" {
		fmt.Fprintln(os.Stderr, "error: input file path (-in) is required")
		os.Exit(1)
	}

	var reader io.Reader
	if *inPath == "-" {
		reader = os.Stdin
	} else {
		file, err := os.Open(*inPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error opening input file '%s': %v\n", *inPath, err)
			os.Exit(1)
		}
		defer file.Close()
		reader = file
	}

	src, err := io.ReadAll(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1)
	}

	src = analysis.NewLexical().Run(src)
	i, ok, err := analysis.NewSyntactical().Run(src)
	if !ok {
		fmt.Println(string(src))
		fmt.Printf("%s^ %s\n", strings.Repeat(" ", i), err.Error())
		fmt.Printf("syntax error at position: %d\n", i)
		os.Exit(1)
	}

	err = compile.New().Run(src, *outPath)
	if err != nil {
		fmt.Println(err)
	}
}

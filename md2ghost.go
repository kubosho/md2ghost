package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jessevdk/go-flags"
)

const version = "v0.0.1"

type opts struct {
	Help    bool   `short:"h" long:"help" description:"Print this message and quit"`
	Version bool   `short:"v" long:"version" description:"Print version information and quit"`
	OutDir  string `short:"o" long:"output" description:"Specify an output directory for Ghost posts"`
}

func main() {
	opts := &opts{}
	p := flags.NewParser(opts, flags.PrintErrors)
	args, err := p.Parse()
	if err != nil {
		return
	}

	if opts.Help || (len(args) == 0 && len(os.Args) < 2) {
		fmt.Fprintf(os.Stderr, helpText)
		return
	}

	if opts.Version {
		fmt.Fprintf(os.Stderr, "md2ghost: %s\n", version)
		return
	}

	var fp *os.File
	var pattern string

	// MEMO: args[0] != "" だと panic: runtime error: index out of range
	if len(args) != 0 {
		pattern = args[0] + "/*.md"
	} else {
		pattern = "." + "/*.md"
	}

	files, err := filepath.Glob(pattern)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	for _, file := range files {
		fp, err = os.Open(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		defer fp.Close()

		scanner := bufio.NewScanner(fp)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	}
}

const helpText = `md2ghost - Convert a markdown files into Ghost posts.

Usage: md2ghost [option] <file|directory>

Options:

  -o, --output  Specify an output directory for Ghost posts
  -h, --help    Print this message and quit
  -v, --version Print version information and quit

Example:

  $ md2ghost .
  $ md2ghost -o path/to/output_directory path/to/your_directory
`

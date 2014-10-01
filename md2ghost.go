package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
)

const version = "v0.0.1"

type opts struct {
	Help    bool   `short:"h" long:"help" description:"Print this message and quit"`
	Version bool   `short:"v" long:"version" description:"Print version information and quit"`
	OutDir  string `short:"o" long:"output" description:"Specify an output directory for Ghost posts"`
}

func init() {
}

func main() {
	opts := &opts{}
	p := flags.NewParser(opts, flags.PrintErrors)
	args, err := p.Parse()
	if err != nil {
		fmt.Fprintf(os.Stderr, helpText)
		return
	}

	if opts.Help {
		fmt.Fprintf(os.Stderr, helpText)
		return
	}

	if opts.Version {
		fmt.Fprintf(os.Stderr, "md2ghost: %s\n", version)
		return
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

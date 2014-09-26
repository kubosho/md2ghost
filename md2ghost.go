package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	flHelp bool
)

func showHelp() {
	fmt.Fprintf(os.Stderr, helpText)
}

func main() {
	flag.BoolVar(&flHelp, "h", false, "Print this message and quit")
	flag.BoolVar(&flHelp, "help", false, "Print this message and quit")
	flag.Parse()

	if flHelp {
		showHelp()
		os.Exit(0)
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

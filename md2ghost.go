package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	flHelp = flag.Bool("help", false, "Print this message and quit")
)

func showHelp() {
	fmt.Fprintf(os.Stderr, helpText)
}

func init() {
	flag.BoolVar(flHelp, "h", false, "Print this message and quit")
}

func main() {
	flag.Parse()

	if *flHelp {
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

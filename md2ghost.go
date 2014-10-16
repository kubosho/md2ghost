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

type GhostJSON struct {
	Meta Meta `json:"meta"`
	Data Data `json:"data"`
}

type Meta struct {
	ExportedOn int64  `json:"exported_on"`
	Version    string `json:"version"`
}

type Data struct {
	Posts []struct {
		AuthorID        int64       `json:"author_id"`
		CreatedAt       int64       `json:"created_at"`
		CreatedBy       int64       `json:"created_by"`
		Featured        int64       `json:"featured"`
		HTML            string      `json:"html"`
		ID              int64       `json:"id"`
		Image           interface{} `json:"image"`
		Language        string      `json:"language"`
		Markdown        string      `json:"markdown"`
		MetaDescription interface{} `json:"meta_description"`
		MetaTitle       interface{} `json:"meta_title"`
		Page            int64       `json:"page"`
		PublishedAt     int64       `json:"published_at"`
		PublishedBy     int64       `json:"published_by"`
		Slug            string      `json:"slug"`
		Status          string      `json:"status"`
		Title           string      `json:"title"`
		UpdatedAt       int64       `json:"updated_at"`
		UpdatedBy       int64       `json:"updated_by"`
	} `json:"posts"`
	PostsTags []struct {
		PostID int64 `json:"post_id"`
		TagID  int64 `json:"tag_id"`
	} `json:"posts_tags"`
	RolesUsers []struct {
		RoleID int64 `json:"role_id"`
		UserID int64 `json:"user_id"`
	} `json:"roles_users"`
	Tags []struct {
		Description string `json:"description"`
		ID          int64  `json:"id"`
		Name        string `json:"name"`
		Slug        string `json:"slug"`
	} `json:"tags"`
	Users []struct {
		Accessibility   interface{} `json:"accessibility"`
		Bio             interface{} `json:"bio"`
		Cover           interface{} `json:"cover"`
		CreatedAt       int64       `json:"created_at"`
		CreatedBy       int64       `json:"created_by"`
		Email           string      `json:"email"`
		ID              int64       `json:"id"`
		Image           interface{} `json:"image"`
		Language        string      `json:"language"`
		LastLogin       interface{} `json:"last_login"`
		Location        interface{} `json:"location"`
		MetaDescription interface{} `json:"meta_description"`
		MetaTitle       interface{} `json:"meta_title"`
		Name            string      `json:"name"`
		Slug            string      `json:"slug"`
		Status          string      `json:"status"`
		UpdatedAt       int64       `json:"updated_at"`
		UpdatedBy       int64       `json:"updated_by"`
		Website         interface{} `json:"website"`
	} `json:"users"`
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

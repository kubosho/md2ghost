md2ghost
========

Convert a markdown files into Ghost posts.

## Description

A tool to convert from Markdown files to Ghost posts didn't until now. So I made this tool a convert the Ghost posts from Markdown files.

## Demo

## Usage

Convert *.md to Ghost posts in a current directory:

```
md2ghost .
```

Specify a source directory where are located *.md:

```
md2ghost path/to/your_directory
```

## Options

Specify an output directory for Ghost posts (default the output is current directory):

```
md2ghost -o path/to/output_directory path/to/your_directory
```

## Install

```
go get github.com/kubosho/md2ghost
```

## Contribution

## Licence

MIT

## Author

[kubosho](https://github.com/kubosho)

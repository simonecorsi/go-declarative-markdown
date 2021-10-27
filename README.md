# Declarative Markdown

<p align="center"><img src="https://raw.githubusercontent.com/simonecorsi/go-declarative-markdown/main/logo.png" height="350px" alt="Declarative Markdown Gopher Logo"/></p>

## About

This package helps in generating markdown in a declarative way without having to handle strings yourself

> This README has been generated using this own package!
You can see the code example [here](./docs/generator.go)

## Requirements

- git
- go 1.16

## Installation

```
go get github.com/simonecorsi/go-declarative-markdown
```

## Usage

```
package main

import mkd "github.com/simonecorsi/go-declarative-markdown/markdown"

function main() {
  md := mkd.CreateMarkdown("My h1 Header")
  md.Header("About")
  md.Paragraph("Lorem Ipsum")
  md.Render() // save to file or whatever
}

```

## API

### CreateMarkdown

`CreateMarkdown(title string) *Markdown`

Initialize new Markdown instance with h1 title on top.

> Returns reference to the Markdown to allow chaining

### GetLine

`GetLine(lineNumber int) string`

Utilities to get a line in the markdown, this is a slice before rendering so LineBreak don't counts

### Render

`Render() string`

Renders all the content add in this markdown instance

### Header

`Header(text string, weigth int) *Markdown`

Creates an heading, you can specify the weigth [1-6]

> Returns reference to the Markdown to allow chaining

### AddLine

`AddLine(text string) *Markdown`

Utilities that pushes content to the slices of lines, this is used internally by each other command

> Returns reference to the Markdown to allow chaining

### Paragraph

`Paragraph(text string) *Markdown`

Create a generic paragraph of text

> Returns reference to the Markdown to allow chaining

### Quote

`Quote(text string) *Markdown`

Create quoted paragraph of text

> Returns reference to the Markdown to allow chaining

### Code

`Code(text string, language string) *Markdown`

Creates a snippet of code with syntax highligh

> Returns reference to the Markdown to allow chaining

### HorizontalLine

`HorizontalLine() *Markdown`

Creates an horizontal line to dive sections

> Returns reference to the Markdown to allow chaining

### Image

`Image(altText string, filepath string) *Markdown`

Creates an image with a title

> Returns reference to the Markdown to allow chaining

### List

`List(items []ListItem, numbered bool) *Markdown`

Creates a list of item, either numbered or with dashes, default to dashes

> Returns reference to the Markdown to allow chaining

```
type ListItem struct {
	Label string
	Depth int
}
```

### Task

`Task(items []TaskItem, numbered bool) *Markdown`

Creates a list with checkboxes

> Returns reference to the Markdown to allow chaining

```
type TaskItem struct {
	Label   string
	Checked bool
}
```

### Table

`Table(headers []string, rows [][]string) *Markdown`

Creates a table

> Returns reference to the Markdown to allow chaining

## Contributing

Project is pretty simple and straight forward for what is my needs, but if you have any idea you're welcome.

This projects uses [Conventional Commit Format](https://www.conventionalcommits.org) so be sure to use standard commit format or PR won't be accepted

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

Simone Corsi - [@im_simonecorsi](https://twitter.com/im_simonecorsi)
package main

import (
	"fmt"
	"os"

	mkd "github.com/simonecorsi/go-declarative-markdown/markdown"
)

func GenerateReadme() *mkd.Markdown {
	md := mkd.CreateMarkdown("Declarative Markdown")

	md.AddLine("<p align=\"center\"><img src=\"https://raw.githubusercontent.com/simonecorsi/go-declarative-markdown/main/banner.png\" alt=\"Declarative Markdown Gopher Logo\"/></p>")

	md.Header("About", 2)
	md.Paragraph("This package helps in generating markdown in a declarative way without having to handle strings yourself")
	md.Quote("This README has been generated using this own package!\nYou can see the code example " + mkd.Link("here", "./docs/generator.go"))

	md.Header("Requirements", 2)
	requirements := append(
		make([]mkd.ListItem, 0),
		mkd.ListItem{Label: "git", Depth: 0},
		mkd.ListItem{Label: "go 1.16", Depth: 0},
	)
	md.List(requirements, false)

	md.Header("Installation", 2)
	md.Code("go get github.com/simonecorsi/go-declarative-markdown", "go")

	md.Header("Usage", 2)
	code := ""
	code += "package main\n\n"
	code += "import mkd \"github.com/simonecorsi/go-declarative-markdown/markdown\"\n\n"
	code += "function main() {\n"
	code += "  md := mkd.CreateMarkdown(\"My h1 Header\")\n"
	code += "  md.Header(\"About\")\n"
	code += "  md.Paragraph(\"Lorem Ipsum\")\n"
	code += "  md.Render() // save to file or whatever\n"
	code += "}\n"
	md.Code(code, "go")

	md.Header("API", 2)

	md.Header("CreateMarkdown", 3)
	md.Paragraph(mkd.InlineCode("CreateMarkdown(title string) *Markdown"))
	md.Paragraph("Initialize new Markdown instance with h1 title on top.")
	md.Quote("Returns reference to the Markdown to allow chaining")

	md.Header("GetLine", 3)
	md.Paragraph(mkd.InlineCode("GetLine(lineNumber int) string"))
	md.Paragraph("Utilities to get a line in the markdown, this is a slice before rendering so LineBreak don't counts")

	md.Header("Render", 3)
	md.Paragraph(mkd.InlineCode("Render() string"))
	md.Paragraph("Renders all the content add in this markdown instance")

	md.Header("Header", 3)
	md.Paragraph(mkd.InlineCode("Header(text string, weigth int) *Markdown"))
	md.Paragraph("Creates an heading, you can specify the weigth [1-6]")
	md.Quote("Returns reference to the Markdown to allow chaining")

	md.Header("AddLine", 3)
	md.Paragraph(mkd.InlineCode("AddLine(text string) *Markdown"))
	md.Paragraph("Utilities that pushes content to the slices of lines, this is used internally by each other command")
	md.Quote("Returns reference to the Markdown to allow chaining")

	md.Header("Paragraph", 3)
	md.Paragraph(mkd.InlineCode("Paragraph(text string) *Markdown"))
	md.Paragraph("Create a generic paragraph of text")
	md.Quote("Returns reference to the Markdown to allow chaining")

	md.Header("Quote", 3)
	md.Paragraph(mkd.InlineCode("Quote(text string) *Markdown"))
	md.Paragraph("Create quoted paragraph of text")
	md.Quote("Returns reference to the Markdown to allow chaining")

	md.Header("Code", 3)
	md.Paragraph(mkd.InlineCode("Code(text string, language string) *Markdown"))
	md.Paragraph("Creates a snippet of code with syntax highligh")
	md.Quote("Returns reference to the Markdown to allow chaining")

	md.Header("HorizontalLine", 3)
	md.Paragraph(mkd.InlineCode("HorizontalLine() *Markdown"))
	md.Paragraph("Creates an horizontal line to dive sections")
	md.Quote("Returns reference to the Markdown to allow chaining")

	md.Header("Image", 3)
	md.Paragraph(mkd.InlineCode("Image(altText string, filepath string) *Markdown"))
	md.Paragraph("Creates an image with a title")
	md.Quote("Returns reference to the Markdown to allow chaining")

	md.Header("List", 3)
	md.Paragraph(mkd.InlineCode("List(items []ListItem, numbered bool) *Markdown"))
	md.Paragraph("Creates a list of item, either numbered or with dashes, default to dashes")
	md.Quote("Returns reference to the Markdown to allow chaining")
	md.Code("type ListItem struct {\n\tLabel string\n\tDepth int\n}", "go")

	md.Header("Task", 3)
	md.Paragraph(mkd.InlineCode("Task(items []TaskItem, numbered bool) *Markdown"))
	md.Paragraph("Creates a list with checkboxes")
	md.Quote("Returns reference to the Markdown to allow chaining")
	md.Code("type TaskItem struct {\n\tLabel   string\n\tChecked bool\n}", "go")

	md.Header("Table", 3)
	md.Paragraph(mkd.InlineCode("Table(headers []string, rows [][]string) *Markdown"))
	md.Paragraph("Creates a table")
	md.Quote("Returns reference to the Markdown to allow chaining")

	md.Header("Contributing", 2)
	md.Paragraph("Project is pretty simple and straight forward for what is my needs, but if you have any idea you're welcome.")
	md.Paragraph(fmt.Sprintf("This projects uses %s so be sure to use standard commit format or PR won't be accepted", mkd.Link("Conventional Commit Format", "https://www.conventionalcommits.org")))

	md.Header("License", 2)
	md.Paragraph("Distributed under the MIT License. See `LICENSE` for more information.")

	md.Header("Contact", 2)
	md.Paragraph("Simone Corsi - " + mkd.Link("@im_simonecorsi", "https://twitter.com/im_simonecorsi"))

	return md
}

func main() {
	os.WriteFile("README.md", []byte(GenerateReadme().Render()), 0666)
}

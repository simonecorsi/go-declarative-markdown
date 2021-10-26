package main

import (
	"fmt"

	mkd "github.com/simonecorsi/go-declarative-markdown"
)

func GenerateReadme() *mkd.Markdown {
	md := mkd.CreateMarkdown("Declarative Markdown")

	md.Header("About", 2)
	md.Paragraph("This package helps in generating markdown in a declarative way without having to handle strings yourself")

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
	code += "import mkd \"github.com/simonecorsi/go-declarative-markdown\"\n\n"
	code += "function main() {\n"
	code += "  md := mkd.CreateMarkdown(\"My h1 Header\")\n"
	code += "  md.Header(\"About\")\n"
	code += "  md.Paragraph(\"Lorem Ipsum\")\n"
	code += "  md.Render() // save to file or whatever\n"
	code += "}\n"
	md.Code(code, "go")

	md.Header("API", 2)

	md.Header("CreateMarkdown", 3)
	md.Paragraph("`CreateMarkdown(title string) *Markdown`")
	md.Header("GetLine", 3)
	md.Header("Render", 3)
	md.Header("Header", 3)
	md.Header("AddLine", 3)
	md.Header("Paragraph", 3)
	md.Header("Quote", 3)
	md.Header("Code", 3)
	md.Header("Image", 3)
	md.Header("List", 3)
	md.Header("Task", 3)
	md.Header("Table", 3)

	return md
}

func main() {
	fmt.Println(GenerateReadme().Render())
}

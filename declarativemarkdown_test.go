package declarativemarkdown

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	md := CreateMarkdown("title")
	if fmt.Sprintf("%T", md.data) != "[]string" {
		t.Errorf("Initialization didn't add default values")
	}
}

func TestHeader(t *testing.T) {
	md := CreateMarkdown("Title")

	if md.data[0] != "# Title" {
		t.Errorf("Header method failed")
	}

	md.Header("MyHeader", 2)
	if md.data[1] != "## MyHeader" {
		t.Errorf("Header method failed")
	}

	md.Header("MyHeader", 3)
	if md.data[2] != "### MyHeader" {
		t.Errorf("Header method failed")
	}

	md.Header("MyHeader", 0)
	if md.data[3] != "# MyHeader" {
		t.Errorf("Didnt default to 1")
	}

	md.Header("MyHeader", 7)
	if md.data[4] != "###### MyHeader" {
		t.Errorf("Didnt default to 6")
	}
}

func TestRender(t *testing.T) {
	md := CreateMarkdown("Title")
	md.Header("MyHeader", 2)
	if md.Render() != fmt.Sprintf("# Title%s## MyHeader%s", LineBreak, LineBreak) {
		t.Errorf("Rendered content differs")
	}
}

// func TestGetLine(t *testing.T) {
// 	md := CreateMarkdown("Title")
// 	if md.GetLine(0) != fmt.Sprintf("Title") {
// 		t.Errorf("Paragraph is different from input")
// 	}
// }

func TestParagraph(t *testing.T) {
	md := CreateMarkdown("Title")
	txt := "Lorem Ipsum Ajeje Brazorf"
	md.Paragraph(txt)
	if md.GetLine(1) != fmt.Sprintf("Lorem Ipsum Ajeje Brazorf%s", LineBreak) {
		t.Errorf("Paragraph is different from input")
	}
}

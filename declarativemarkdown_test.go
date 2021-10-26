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
	txt := "MyHeader"
	md := CreateMarkdown(txt)

	if md.GetLine(0) != fmt.Sprintf("# %s", txt) {
		t.Errorf("Header method failed")
	}

	md.Header(txt, 2)
	if md.GetLine(1) != fmt.Sprintf("## %s", txt) {
		t.Errorf("Header method failed")
	}

	md.Header(txt, 3)
	if md.GetLine(2) != fmt.Sprintf("### %s", txt) {
		t.Errorf("Header method failed")
	}

	md.Header(txt, 0)
	if md.GetLine(3) != fmt.Sprintf("# %s", txt) {
		t.Errorf("Didnt default to 1")
	}

	md.Header(txt, 7)
	if md.GetLine(4) != fmt.Sprintf("###### %s", txt) {
		t.Errorf("Didnt default to 6")
	}
}

func TestRender(t *testing.T) {
	md := CreateMarkdown("Title")
	if md.Render() != "# Title" {
		t.Errorf("Rendered content differs")
	}
}

func TestGetLine(t *testing.T) {
	md := CreateMarkdown("Title")
	if md.GetLine(0) != "# Title" {
		t.Errorf("Paragraph is different from input")
	}
}

func TestParagraph(t *testing.T) {
	md := CreateMarkdown("Title")
	txt := "Lorem Ipsum Ajeje Brazorf"
	md.Paragraph(txt)
	if md.GetLine(1) != txt {
		t.Errorf("Paragraph is different from input")
	}
}

func TestImage(t *testing.T) {
	md := CreateMarkdown("Title")
	txt := "here"
	img := "./image.png"
	md.Image(txt, img)
	if md.GetLine(1) != fmt.Sprintf("![%s](%s)", txt, img) {
		t.Errorf("Paragraph is different from input")
	}
}

func TestList(t *testing.T) {
	md := CreateMarkdown("Title")
	list := make([]ListItem, 0)
	list = append(list, ListItem{label: "ajeje", depth: 0}, ListItem{label: "brazorf", depth: 1})
	md.List(list, false)
	out := fmt.Sprintf("- ajeje%s  - brazorf", LineBreak)
	if md.GetLine(1) != out {
		t.Errorf("List is different from input")
	}
}

func TestTask(t *testing.T) {
	md := CreateMarkdown("Title")
	task := make([]TaskItem, 0)
	task = append(task, TaskItem{label: "ajeje", checked: true}, TaskItem{label: "brazorf", checked: false})
	md.Task(task, false)
	out := fmt.Sprintf("[X] ajeje%s[ ] brazorf", LineBreak)
	if md.GetLine(1) != out {
		t.Errorf("Task is different from input")
	}
}

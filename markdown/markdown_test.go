package markdown

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	md := CreateMarkdown("Init")
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
	md := CreateMarkdown("Render")
	if md.Render() != "# Render" {
		t.Errorf("Rendered content differs")
	}
}

func TestGetLine(t *testing.T) {
	md := CreateMarkdown("GetLine")
	if md.GetLine(0) != "# GetLine" {
		t.Errorf("Paragraph is different from input")
	}
}

func TestParagraph(t *testing.T) {
	md := CreateMarkdown("Paragraph")
	txt := "Lorem Ipsum Ajeje Brazorf"
	md.Paragraph(txt)
	if md.GetLine(1) != txt {
		t.Errorf("Paragraph is different from input")
	}
}

func TestCode(t *testing.T) {
	md := CreateMarkdown("Code")
	txt := "package main\nfunction main() {\n}"
	md.Code(txt, "go")
	if md.GetLine(1) != fmt.Sprintf("```\n%s\n```", txt) {
		t.Errorf("Code is different from input")
	}
}

func TestQuote(t *testing.T) {
	md := CreateMarkdown("Quote")
	txt := "Lorem Ipsum Ajeje Brazorf"
	md.Quote(txt)
	if md.GetLine(1) != fmt.Sprintf("> %s", txt) {
		t.Errorf("Quote is different from input")
	}
}

func TestHorizontalLine(t *testing.T) {
	md := CreateMarkdown("HorizontalLine")
	txt := "---"
	md.HorizontalLine()
	if md.GetLine(1) != txt {
		t.Errorf("HorizontalLine is different from input")
	}
}

func TestImage(t *testing.T) {
	md := CreateMarkdown("Image")
	txt := "here"
	img := "./image.png"
	md.Image(txt, img)
	if md.GetLine(1) != fmt.Sprintf("![%s](%s)", txt, img) {
		t.Errorf("Image is different from input")
	}
}

func TestList(t *testing.T) {
	md := CreateMarkdown("List")
	list := make([]ListItem, 0)
	list = append(list, ListItem{Label: "ajeje", Depth: 0}, ListItem{Label: "brazorf", Depth: 1})
	md.List(list, false)
	out := fmt.Sprintf("- ajeje%s  - brazorf", LineBreak)
	if md.GetLine(1) != out {
		t.Errorf("List is different from input")
	}
}

func TestTask(t *testing.T) {
	md := CreateMarkdown("Task")
	task := make([]TaskItem, 0)
	task = append(task, TaskItem{Label: "ajeje", Checked: true}, TaskItem{Label: "brazorf", Checked: false})
	md.Task(task, false)
	out := fmt.Sprintf("[X] ajeje%s[ ] brazorf", LineBreak)
	if md.GetLine(1) != out {
		t.Errorf("Task is different from input")
	}
}

func TestTable(t *testing.T) {
	md := CreateMarkdown("Table")

	headers := make([]string, 0)
	headers = append(headers, "id", "username")

	row := make([]string, 0)
	row = append(row, "1", "ajeje_brazorf")

	rows := make([][]string, 0)
	rows = append(rows, row)

	md.Table(headers, rows)

	out := "| id | username |"
	out += LineBreak + "| --- | --- |"
	out += LineBreak + "| 1 | ajeje_brazorf |"

	if md.GetLine(1) != out {
		t.Errorf("Table is different from input")
	}
}

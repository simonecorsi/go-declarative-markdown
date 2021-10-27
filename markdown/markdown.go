package markdown

import (
	"fmt"
	"strings"
)

const LineBreak = "\r\n"

type Markdown struct {
	data []string
}

func CreateMarkdown(title string) *Markdown {
	data := make([]string, 0)

	md := &Markdown{
		data: data,
	}

	md.Header(title, 1)
	return md
}

func (m *Markdown) GetLine(lineNumber int) string {
	return m.data[lineNumber]
}

func (m *Markdown) Render() string {
	return strings.Join(m.data, LineBreak+LineBreak)
}

func (m *Markdown) Header(text string, weigth int) *Markdown {
	if weigth > 6 {
		weigth = 6
	}
	if weigth < 1 {
		weigth = 1
	}
	m.AddLine(fmt.Sprintf("%s %s", strings.Repeat("#", weigth), text))
	return m
}

func (m *Markdown) AddLine(text string) *Markdown {
	m.data = append(m.data, text)
	return m
}

func (m *Markdown) Paragraph(text string) *Markdown {
	m.AddLine(text)
	return m
}

func (m *Markdown) Quote(text string) *Markdown {
	m.AddLine(fmt.Sprintf("> %s", text))
	return m
}

func (m *Markdown) Code(text string, language string) *Markdown {
	m.AddLine(fmt.Sprintf("```\n%s\n```", text))
	return m
}

func (m *Markdown) HorizontalLine() *Markdown {
	m.AddLine("---")
	return m
}

func (m *Markdown) Image(altText string, filepath string) *Markdown {
	m.AddLine(fmt.Sprintf("![%s](%s)", altText, filepath))
	return m
}

type ListItem struct {
	Label string
	Depth int
}

func (m *Markdown) List(items []ListItem, numbered bool) *Markdown {
	bullet := "-"
	var results []string
	for i, item := range items {
		if numbered {
			bullet = fmt.Sprintf("%d.", i+1)
		}
		results = append(results, fmt.Sprintf(
			"%s%s %s",
			// spaces
			strings.Repeat("  ", item.Depth),
			// bullet type
			bullet,
			// Label
			item.Label,
		))
	}

	m.AddLine(strings.Join(results, LineBreak))
	return m
}

type TaskItem struct {
	Label   string
	Checked bool
}

func (m *Markdown) Task(items []TaskItem, numbered bool) *Markdown {
	var results []string
	for _, task := range items {
		check := " "
		if task.Checked {
			check = "X"
		}
		results = append(results, fmt.Sprintf(
			"[%s] %s",
			check,
			task.Label,
		))
	}

	m.AddLine(strings.Join(results, LineBreak))
	return m
}

func (m *Markdown) Table(headers []string, rows [][]string) *Markdown {
	t := ""

	t += fmt.Sprintf("| %s |", strings.Join(headers, " | "))
	t += fmt.Sprintf("%s|%s", LineBreak, strings.Repeat(" --- |", len(headers)))

	for _, row := range rows {
		t += fmt.Sprintf("%s| %s |", LineBreak, strings.Join(row, " | "))
	}

	m.AddLine(t)
	return m
}

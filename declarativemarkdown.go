package declarativemarkdown

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

func (m *Markdown) GetLine(line int) string {
	return m.data[line]
}

func (m *Markdown) Render() string {
	return strings.Join(m.data, LineBreak)
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
	m.data = append(m.data, fmt.Sprintf("%s%s", text, LineBreak))
	return m
}

func (m *Markdown) Paragraph(text string) *Markdown {
	m.AddLine(text)
	return m
}

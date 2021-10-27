package markdown

import "fmt"

func Italic(text string) string {
	return fmt.Sprintf("*%s*", text)
}

func Bold(text string) string {
	return fmt.Sprintf("**%s**", text)
}

func Strike(text string) string {
	return fmt.Sprintf("~~%s~~", text)
}

func InlineCode(text string) string {
	return fmt.Sprintf("`%s`", text)
}

package markdown

import (
	"fmt"
	"testing"
)

func TestItalic(t *testing.T) {
	txt := "My Test"
	out := Italic(txt)
	expect := fmt.Sprintf("*%s*", txt)
	if expect != out {
		t.Errorf("[Invalid output]\nExpected:\n%s\n\nGot:\n%s\n", expect, out)
	}
}

func TestBold(t *testing.T) {
	txt := "My Test"
	out := Bold(txt)
	expect := fmt.Sprintf("**%s**", txt)
	if expect != out {
		t.Errorf("[Invalid output]\nExpected:\n%s\n\nGot:\n%s\n", expect, out)
	}
}

func TestStrike(t *testing.T) {
	txt := "My Test"
	out := Strike(txt)
	expect := fmt.Sprintf("~~%s~~", txt)
	if expect != out {
		t.Errorf("[Invalid output]\nExpected:\n%s\n\nGot:\n%s\n", expect, out)
	}
}

func TestLink(t *testing.T) {
	txt := "My Test"
	link := "http://google.com"
	out := Link(txt, link)
	expect := fmt.Sprintf("[%s](%s)", txt, link)
	if expect != out {
		t.Errorf("[Invalid output]\nExpected:\n%s\n\nGot:\n%s\n", expect, out)
	}
}

func TestInlineCode(t *testing.T) {
	txt := "My Test"
	out := InlineCode(txt)
	expect := fmt.Sprintf("`%s`", txt)
	if expect != out {
		t.Errorf("[Invalid output]\nExpected:\n%s\n\nGot:\n%s\n", expect, out)
	}
}

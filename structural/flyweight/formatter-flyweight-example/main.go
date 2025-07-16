package main

import (
	"fmt"
	"strings"
	"unicode"
)

// It is a anti-flyweight example
type FormattedText struct {
	plainText  string
	capitalize []bool // same size of plainText
}

func (f *FormattedText) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(f.plainText); i++ {
		c := f.plainText[i]
		if f.capitalize[i] {
			sb.WriteRune(unicode.ToUpper(rune(c)))
		} else {
			sb.WriteRune(rune(c))
		}
	}
	return sb.String()
}

func (f *FormattedText) Capitalize(start, end int) {
	for i := start; i <= end; i++ {
		f.capitalize[i] = true
	}
}

func NewFormattedText(plainText string) *FormattedText {
	return &FormattedText{plainText: plainText, capitalize: make([]bool, len(plainText))}
}

// It is a Flyweight
type TextRange struct {
	Start, End               int
	Capitalize, Bold, Italic bool
}

func (t *TextRange) Covers(position int) bool {
	return position >= t.Start && position <= t.End
}

type BetterFormattedText struct {
	plainText  string
	formatting []*TextRange
}

func NewBetterFormattedText(plainText string) *BetterFormattedText {
	return &BetterFormattedText{plainText: plainText}
}

// Fluent function
func (b *BetterFormattedText) Range(start, end int) *TextRange {
	r := &TextRange{start, end, false, false, false}
	b.formatting = append(b.formatting, r)
	return r
}

func (b *BetterFormattedText) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(b.plainText); i++ {
		c := b.plainText[i]
		for _, r := range b.formatting {
			if r.Covers(i) && r.Capitalize {
				c = uint8(unicode.ToUpper(rune(c)))
			}
		}
		sb.WriteRune(rune(c))
	}
	return sb.String()
}

func main() {
	text := "The new world order"

	// Bad implementation
	ft := NewFormattedText(text)
	ft.Capitalize(8, 12)
	fmt.Println(ft.String())

	// FLyweight implementation
	bft := NewBetterFormattedText(text)
	bft.Range(8, 12).Capitalize = true
	fmt.Println(bft.String())
}

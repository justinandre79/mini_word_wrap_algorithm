// Package mini_word_wrap_algorithm breaks text into lines of a maximum width
// using a greedy wrapping strategy.
//
// The greedy strategy is deliberate: it is simple, predictable, and linear in
// the size of the input. It does not attempt to balance line lengths or avoid
// rivers. If you need optimal-fit wrapping, this package is the wrong tool.
package mini_word_wrap_algorithm

import (
	"strings"
	"unicode"
)

// Wrap greedily wraps text into lines no wider than width runes.
//
// Words are separated by runs of Unicode whitespace. A word that is longer
// than width is placed on its own line and is not broken; this keeps the
// function honest about what "maximum width" means when a single token cannot
// fit. Leading and trailing whitespace on each line is discarded, and runs of
// whitespace between words collapse to a single space in the output.
//
// Width is measured in runes, not bytes, so multi-byte characters count as one
// column. Width must be positive; Wrap panics if width < 1 because there is no
// sensible line to return.
func Wrap(text string, width int) []string {
	if width < 1 {
		panic("mini_word_wrap_algorithm: width must be at least 1")
	}

	words := strings.FieldsFunc(text, unicode.IsSpace)
	if len(words) == 0 {
		return nil
	}

	lines := make([]string, 0, 1+len(text)/width)
	current := ""
	currentWidth := 0

	for _, word := range words {
		wordWidth := len([]rune(word))

		if currentWidth == 0 {
			current = word
			currentWidth = wordWidth
			continue
		}

		if currentWidth+1+wordWidth <= width {
			current += " " + word
			currentWidth += 1 + wordWidth
			continue
		}

		lines = append(lines, current)
		current = word
		currentWidth = wordWidth
	}

	lines = append(lines, current)
	return lines
}

// WrapString is a convenience wrapper around Wrap that joins the resulting
// lines with "\n". It returns the empty string for empty input.
func WrapString(text string, width int) string {
	lines := Wrap(text, width)
	if len(lines) == 0 {
		return ""
	}
	return strings.Join(lines, "\n")
}

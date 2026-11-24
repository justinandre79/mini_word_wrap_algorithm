# mini_word_wrap_algorithm

Greedy word wrapping: split a string into lines no wider than a given number of runes.

## Usage

```go
package main

import (
	"fmt"

	"mini_word_wrap_algorithm"
)

func main() {
	lines := mini_word_wrap_algorithm.Wrap("the quick brown fox jumps", 10)
	for _, line := range lines {
		fmt.Println(line)
	}
	// the quick
	// brown fox
	// jumps

	fmt.Println(mini_word_wrap_algorithm.WrapString("the quick brown fox jumps", 10))
	// the quick
	// brown fox
	// jumps
}
```

`Wrap` returns `[]string`; `WrapString` joins the same lines with `"\n"`.

## Why this exists

Line breaking is a recurring chore in terminal UIs, log formatters, and text
previews. A greedy algorithm is the smallest thing that works: scan words left
to right, append while the next word fits, otherwise start a new line. It is
linear in the input and has no configuration surface. The trade-off is that it
does not balance line lengths the way an optimal-fit (Knuth-Plass) algorithm
would; short last lines and uneven right edges are expected.

## The awkward edge

A single word longer than `width` is emitted on its own line and is *not*
broken. That is the only honest behaviour when a token cannot fit: breaking it
would change the text, and dropping it would lose data. Callers that need
character-level breaking should post-process the returned lines. Width is
counted in runes, so multi-byte characters count as one column; `Wrap` panics
if `width` is less than 1.

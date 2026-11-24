package mini_word_wrap_algorithm

import (
	"reflect"
	"testing"
)

func TestWrapBasic(t *testing.T) {
	got := Wrap("the quick brown fox jumps over the lazy dog", 10)
	want := []string{
		"the quick",
		"brown fox",
		"jumps over",
		"the lazy",
		"dog",
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Wrap = %#v, want %#v", got, want)
	}
}

func TestWrapEmptyInput(t *testing.T) {
	got := Wrap("", 10)
	if got != nil {
		t.Fatalf("Wrap(\"\", 10) = %#v, want nil", got)
	}
}

func TestWrapWhitespaceOnlyInput(t *testing.T) {
	got := Wrap("   \t\n  ", 10)
	if got != nil {
		t.Fatalf("Wrap(whitespace, 10) = %#v, want nil", got)
	}
}

func TestWrapCollapsesInternalWhitespace(t *testing.T) {
	got := Wrap("one\t\ttwo   three", 20)
	want := []string{"one two three"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Wrap = %#v, want %#v", got, want)
	}
}

func TestWrapExactFit(t *testing.T) {
	got := Wrap("ab cd", 5)
	want := []string{"ab cd"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Wrap = %#v, want %#v", got, want)
	}
}

func TestWrapOneRuneOverflows(t *testing.T) {
	got := Wrap("ab cd", 4)
	want := []string{"ab", "cd"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Wrap = %#v, want %#v", got, want)
	}
}

func TestWrapLongWordNotBroken(t *testing.T) {
	got := Wrap("a supercalifragilistic b", 5)
	want := []string{"a", "supercalifragilistic", "b"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Wrap = %#v, want %#v", got, want)
	}
}

func TestWrapUnicodeCountsRunes(t *testing.T) {
	// Each of these runes is multiple bytes in UTF-8 but one column wide.
	got := Wrap("\u00e9\u00e9\u00e9 \u00e9\u00e9", 4)
	want := []string{"\u00e9\u00e9\u00e9", "\u00e9\u00e9"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Wrap = %#v, want %#v", got, want)
	}
}

func TestWrapWidthOne(t *testing.T) {
	got := Wrap("a b c", 1)
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Wrap = %#v, want %#v", got, want)
	}
}

func TestWrapPanicsOnZeroWidth(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("Wrap(text, 0) did not panic")
		}
	}()
	Wrap("anything", 0)
}

func TestWrapPanicsOnNegativeWidth(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("Wrap(text, -3) did not panic")
		}
	}()
	Wrap("anything", -3)
}

func TestWrapString(t *testing.T) {
	got := WrapString("the quick brown fox", 9)
	want := "the quick\nbrown fox"
	if got != want {
		t.Fatalf("WrapString = %q, want %q", got, want)
	}
}

func TestWrapStringEmpty(t *testing.T) {
	got := WrapString("", 10)
	if got != "" {
		t.Fatalf("WrapString(\"\", 10) = %q, want \"\"", got)
	}
}

package gobyexample

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestString(t *testing.T) {
	const s = "胡军"

	fmt.Println("len:", len(s)) // 6

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println()
	fmt.Println("rune:", utf8.RuneCountInString(s)) // 2

	for idx, runeVal := range s {
		fmt.Printf("%#U starts at %d\n", runeVal, idx)
	}

	fmt.Println("\n Using DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeVale, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeVale, i)
		w = width
		examineRune(runeVale)
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'p' {
		fmt.Println("found so")
	} else {
		fmt.Println("found:", r)
	}
}

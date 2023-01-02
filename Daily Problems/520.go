package main

import (
	"fmt"
	"unicode"
)

func detectCapitalUse(word string) bool {
	last := unicode.IsUpper(rune(word[len(word)-1]))
	for i := len(word) - 1; i > 0; i-- {
		if last != unicode.IsUpper(rune(word[i])) {
			return false
		}
	}

	if last == true && false == unicode.IsUpper(rune(word[0])) {
		return false
	} else {
		return true

	}
}
func main() {
	word := "aaA"
	fmt.Println(detectCapitalUse(word))
}

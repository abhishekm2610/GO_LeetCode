package main

import (
	"fmt"
	"strings"
)

//Original Solution
func mostWordsFound(sentences []string) int {
	maxWords := 0
	for i := 0; i < len(sentences); i++ {
		nextSentWords := len(strings.Split(sentences[i], " "))
		if maxWords < nextSentWords {
			maxWords = nextSentWords
		}
	}
	return maxWords
}

func main() {
	input := []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}

	fmt.Println(mostWordsFound((input)))
}

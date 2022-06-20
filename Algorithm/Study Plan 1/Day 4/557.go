package main

import "fmt"

func reverseString(s []byte) {
	i := 0
	j := len(s) - 1
	for i < j {
		t := s[i]
		s[i] = s[j]
		s[j] = t
		i++
		j--
	}
	return
}

func reverseWords(s string) string {

	startingIndex := 0
	word := []byte{}
	sentence := ""
	for startingIndex < len(s) {
		if s[startingIndex] != ' ' {
			word = append(word, s[startingIndex])
		} else {
			reverseString(word)
			sentence += string(word)
			sentence += " "

			word = []byte{}

		}
		startingIndex++
	}
	reverseString(word)
	sentence += string(word)

	return sentence
}

func main() {
	s := "Let's take LeetCode contest"
	fmt.Println(reverseWords(s))
}

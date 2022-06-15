package main

import (
	"fmt"
)

func countUniqueChar(text string) int {
	var n int = len(text)
	if n == 0 {
		return 0
	}
	var result int = 0
	var record = make([]int, 256)

	for i := 0; i < n; i++ {
		record[text[i]] += 1
	}
	for i := 0; i < n; i++ {
		if record[text[i]] == 1 {

			result++
		}
	}
	return result
}

//Original Solution
func buddyStrings(s string, goal string) bool {

	hashmap := make(map[byte]int)

	for i, _ := range s {
		if s[i] != goal[i] {
			hashmap[s[i]] = i
		}
	}

	if len(hashmap) == 2 {
		keys := []byte{}
		for i, _ := range hashmap {
			keys = append(keys, i)
		}

		s = s[:hashmap[keys[0]]] + string(keys[1]) + s[hashmap[keys[0]]+1:]
		s = s[:hashmap[keys[1]]] + string(keys[0]) + s[hashmap[keys[1]]+1:]

		if s == goal {
			return true
		}
	} else if len(hashmap) == 0 && s == goal && len(s) > countUniqueChar(s) {
		return true

	}

	return false
}
func main() {
	input1 := "ab"
	ip2 := "ab"

	fmt.Println(buddyStrings(input1, ip2))

}

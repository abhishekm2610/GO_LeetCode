package main

import "fmt"

func isSubsequence(s string, t string) bool {
	index := 0
	if len(s) == 0 {
		return true
	} else if len(t) == 0 {
		return false
	}
	for i := range t {
		if t[i] == s[index] {
			index++
			if index == len(s) {
				return true
			}
		}
	}
	return false
}

func main() {
	s, t := "0", "0"

	fmt.Print(s, t, isSubsequence(s, t))
}

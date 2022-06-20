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

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o', '!'}
	reverseString(s)
	fmt.Println(string(s))
}

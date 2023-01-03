package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	dict_s := make(map[byte]byte)
	dict_t := make(map[byte]byte)

	for i := range s {
		v := s[i]
		if _, exists := dict_s[v]; exists && dict_s[v] != t[i] {
			return false
		} else {
			if _, exists := dict_t[t[i]]; exists && dict_t[t[i]] != v {
				return false
			} else {
				dict_t[t[i]] = v
			}
			dict_s[v] = t[i]
		}
	}
	return true
}

func main() {
	s, t := "egg", "add"

	fmt.Print(s, t, isIsomorphic(s, t))
}

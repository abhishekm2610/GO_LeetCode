package main

import "fmt"

type WordFilter struct {
	prefix map[string][]int
	suffix map[string][]int
}

func Constructor(words []string) WordFilter {

	prefixDict := make(map[string][]int)
	for dictIndex, word := range words {
		prefixStr := ""

		for _, v := range word {
			prefixStr += string(v)
			prefixDict[prefixStr] = append(prefixDict[prefixStr], dictIndex)
		}
	}
	suffixDict := make(map[string][]int)
	for dictIndex, word := range words {
		suffixStr := ""

		for i := len(word) - 1; i >= 0; i-- {
			suffixStr = string(word[i]) + suffixStr
			suffixDict[suffixStr] = append(suffixDict[suffixStr], dictIndex)
		}
	}
	wordsObj := WordFilter{prefix: prefixDict, suffix: suffixDict}

	return wordsObj
}

func Intersection(a, b []int) (c []int) {
	m := make(map[int]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

func findMaxElement(arr []int) int {
	max_num := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max_num {
			max_num = arr[i]
		}
	}
	return max_num
}

func (this *WordFilter) F(prefix string, suffix string) int {
	prefixIndices := []int{}
	suffixIndices := make([]int, len(suffix))

	if requiredIndex, isPresent := this.prefix[prefix]; isPresent {
		for _, v := range requiredIndex {
			prefixIndices = append(prefixIndices, v)
		}
	} else {
		return -1
	}
	if requiredIndex, isPresent := this.suffix[suffix]; isPresent {
		for _, v := range requiredIndex {
			suffixIndices = append(suffixIndices, v)
		}

	} else {
		return -1
	}
	fmt.Println(prefixIndices, suffixIndices)
	ans := Intersection(prefixIndices, suffixIndices)
	if len(ans) == 0 {
		return -1
	}
	return findMaxElement(ans)
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(prefix,suffix);
 */
func main() {
	words := []string{"cabaabaaaa", "ccbcababac", "bacaabccba", "bcbbcbacaa", "abcaccbcaa", "accabaccaa", "cabcbbbcca", "ababccabcb", "caccbbcbab", "bccbacbcba"}
	// calls := make([][]string)
	// calls = [,"a"]
	// ,["ab","abcaccbcaa"],["a","aa"],["cabaaba","abaaaa"],["cacc","accbbcbab"],["ccbcab","bac"],["bac","cba"],["ac","accabaccaa"],["bcbb","aa"],["ccbca","cbcababac"]

	// ["bccbacbcba","a"],["ab","abcaccbcaa"],["a","aa"],["cabaaba","abaaaa"],["cacc","accbbcbab"],["ccbcab","bac"],["bac","cba"],["ac","accabaccaa"],["bcbb","aa"],["ccbca","cbcababac"]]
	obj := Constructor(words)
	fmt.Println(obj.prefix)
	fmt.Println(obj.suffix)

	param_1 := obj.F("bccbacbcba", "a")
	fmt.Println(param_1)
}

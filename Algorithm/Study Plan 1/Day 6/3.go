package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {

	ans := 0
	ret := 0
	if len(s) == 1 {
		return 1
	}
	for i := 0; i < len(s)-1; i++ {
		j := i + 1
		ans = 1
		memo := make(map[string]int)
		memo[string(s[i])] = ans
		for j < len(s) {
			if _, isPresent := memo[string(s[j])]; isPresent {
				memo = map[string]int{}
				memo[string(s[j])] = ans
				if ans > ret {
					ret = ans
				}
				break
			} else {
				ans++
				// fmt.Println(string(s[i]), string(s[j]), ans)

				memo[string(s[j])] = ans
			}
			j++
		}
		if ans > ret {
			return ans
		}
	}

	return ret
}

func main() {

	ip1 := "bwf"
	// ip2 := 0
	fmt.Println(lengthOfLongestSubstring(ip1))

}

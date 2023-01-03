package main

import "fmt"

func minDeletionSize(strs []string) int {
	ans := 0

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][i] < strs[j-1][i] {
				ans++
				break
			}
		}
	}
	return ans
}

func main() {
	strs := []string{"zyx", "wvu", "tsr"}

	fmt.Print(strs, minDeletionSize(strs))
}

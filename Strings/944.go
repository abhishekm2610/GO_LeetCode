package main

import (
	"fmt"
)

//Original Solution
func minDeletionSize(strs []string) int {
	deleteCount := 0
	for i := 0; i < len(strs[0]); i++ {
		startingChar := strs[0][i]
		for j := 0; j < len(strs); j++ {
			if startingChar > strs[j][i] {
				deleteCount++
				break
			} else {
				startingChar = strs[j][i]
			}
		}
	}
	return deleteCount
}

//Less Space and Time
func minDeletionSizeOp(strs []string) int {
	deleted := make([]bool, len(strs[0]))
	// fmt.Println(deleted)
	result := 0
	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(deleted); j++ {
			if deleted[j] {
				continue
			}

			if strs[i][j] < strs[i-1][j] {
				deleted[j] = true
				result++
			}
		}
	}

	return result
}

func main() {
	input := []string{"zyx", "wvu", "tsr"}

	fmt.Println(minDeletionSizeOp((input)))
}

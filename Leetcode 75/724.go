package main

import "fmt"

func Sum(n []int) int {
	s := 0
	for _, v := range n {
		s += v
	}
	return s
}

func pivotIndex(nums []int) int {
	rightSum := Sum(nums)
	leftSum := 0

	for i, v := range nums {
		rightSum = rightSum - v
		if leftSum == rightSum {
			return i
		}
		leftSum += v
	}

	return -1
}

func main() {
	nums := []int{2, 1, -1}
	fmt.Print(pivotIndex(nums))

}

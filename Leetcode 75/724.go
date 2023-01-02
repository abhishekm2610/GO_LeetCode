package main

import "fmt"

func pivotindex(nums []int) int {

	pivotSum := 0

	for i := 0; i < len(nums); i++ {
		if i != 0 {
			pivotSum += nums[i-1]
		}
		rightSum := 0
		for j := i + 1; j < len(nums); j++ {
			rightSum += nums[j]
		}
		fmt.Println(pivotSum, rightSum)

		if pivotSum == rightSum {
			return i
		}
	}
	return -1

}

func main() {
	nums := []int{2, 1, -1}
	fmt.Print(pivotindex(nums))

}

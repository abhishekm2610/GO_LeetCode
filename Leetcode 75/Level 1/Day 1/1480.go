package main

import "fmt"

// Order of N (actually N-2)
//Lessons learnt: To use less space by not creating any temps
func runningSum(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + nums[i-1]
		// fmt.Println(nums)
	}
	return nums
}

func main() {
	nums := []int{3, 1, 2, 10, 1}
	fmt.Print(runningSum(nums))
}

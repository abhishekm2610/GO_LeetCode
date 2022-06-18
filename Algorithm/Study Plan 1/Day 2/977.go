package main

import (
	"fmt"
	"math"
)

//Original Solution
func sortedSquares(nums []int) []int {
	i := 0
	j := len(nums) - 1
	res := make([]int, len(nums)) // Copy slice
	copy(res, nums)
	for p := len(nums) - 1; p >= 0; p-- {
		if math.Abs(float64(nums[i])) > math.Abs(float64(nums[j])) {
			res[p] = nums[i] * nums[i]
			i++
		} else {
			res[p] = nums[j] * nums[j]
			j--
		}
	}
	return res
}

func main() {

	ip1 := []int{-7, -3, 2, 3, 11}
	fmt.Println(sortedSquares(ip1))

}

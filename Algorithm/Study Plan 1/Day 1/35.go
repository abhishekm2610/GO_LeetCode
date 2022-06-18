package main

import "fmt"

//Original Solution
func searchInsert(nums []int, target int) int {

	left := 0
	right := len(nums) - 1

	for left <= right {

		mid := left + (right-left)/2

		if nums[mid] == target {

			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {

			right = mid - 1
		}
	}
	return left
}

func main() {

	ip1 := []int{1, 3, 5, 6}
	ip2 := 0
	fmt.Println(searchInsert(ip1, ip2))

}

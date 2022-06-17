package main

import "fmt"

func search(nums []int, target int) int {
	if nums[len(nums)/2] == target {
		return len(nums) / 2
	}
	left := 0
	right := len(nums) - 1
	for left <= right {
		//pivot := right - left - https://ai.googleblog.com/2006/06/extra-extra-read-all-about-it-nearly.html
		pivot := left + (right-left)/2
		if nums[pivot] == target {
			return pivot
		} else if nums[pivot] > target {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}
	return -1
}

func main() {

	ip1 := []int{-1, 0, 3, 5, 9, 1}
	ip2 := 2
	fmt.Println(search(ip1, ip2))

}

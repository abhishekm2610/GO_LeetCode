package main

import "fmt"

//Original Solution
func searchInsert(nums []int, target int) int {
	if nums[len(nums)/2] == target {
		return len(nums) / 2
	}

	left := 0
	right := len(nums) - 1

	for left <= right {
		pivot := left + (right-left)/2

		if nums[pivot] == target {
			return pivot
		}
		if nums[pivot]-1 == target {
			if pivot == 0 {
				return 0
			} else {
				if nums[pivot-1] != target {
					return pivot
				}
			}
		}
		if nums[pivot]+1 == target {
			if pivot == len(nums)-1 {
				return pivot + 1
			} else {
				if nums[pivot+1] != target {
					return pivot + 1
				}
			}

		}

		if nums[pivot] < target {
			left = pivot + 1
		} else {
			right = pivot - 1
		}

	}
	return -1
}

func main() {

	ip1 := []int{1, 3, 5, 6}
	ip2 := 0
	fmt.Println(searchInsert(ip1, ip2))

}

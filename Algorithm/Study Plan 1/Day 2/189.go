package main

import (
	"fmt"
	"reflect"
)

//Original Solution
func reverse(nums []int, l int, r int) {
	i := l
	j := r
	for i < j {
		swapF := reflect.Swapper(nums)
		swapF(i, j)
		j--
		i++
	}

}

func rotate(nums []int, k int) {
	if len(nums) == k || len(nums) == 1 {
		return
	}

	k = k % len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
	// // p := k
	// for j < len(nums) {

	// 	swapF := reflect.Swapper(nums)
	// 	swapF(i, j)
	// 	// if p <= j {
	// 	// 	swapF(p, j)

	// 	// }
	// 	j++
	// 	i++
	// 	// p++
	// }

}
func main() {

	ip1 := []int{1, 2, 3, 4, 5, 6, 7}
	ip2 := 2
	rotate(ip1, ip2)
	fmt.Println(ip1)

}

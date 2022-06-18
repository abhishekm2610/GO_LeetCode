package main

import "fmt"

//helper
func isBadVersion(version int) bool {
	if version == 4 {
		return true
	} else {
		return false
	}
}

//original solution
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
func firstBadVersion(n int) int {

	left := 0
	right := n

	for left <= right {
		pivot := left + (right-left)/2
		fmt.Println(pivot)
		if isBadVersion(pivot) {
			if !isBadVersion(pivot - 1) {
				return pivot
			} else {
				right = pivot
				continue
			}
		} else {
			fmt.Println(pivot)
			if isBadVersion(pivot + 1) {
				return pivot + 1
			} else {
				left = pivot + 1
			}
		}
	}
	return -1
}

func main() {

	ip2 := 2
	fmt.Println(firstBadVersion(ip2))

}

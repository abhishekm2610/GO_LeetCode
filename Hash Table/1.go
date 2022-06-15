package main

import "fmt"

//Original Solution
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, v := range nums {

		if requiredIndex, isPresent := seen[target-v]; isPresent {
			return []int{requiredIndex, i}
		} else {
			seen[v] = i
		}
	}

	return []int{}
}

func main() {
	input1 := []int{3, 3}
	intput2 := 6

	fmt.Println(twoSum(input1, intput2))

}

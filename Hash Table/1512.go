package main

import "fmt"

//Original Solution
func numIdenticalPairs(nums []int) int {
	seen := make(map[int][]int)
	count := 0
	for i, v := range nums {
		if indexes, isPresent := seen[v]; isPresent {
			for j, _ := range indexes {
				if j < i {
					count++
				}
			}
			seen[v] = append(seen[v], i)

		} else {
			seen[v] = []int{i}

		}
	}
	return count
}

func main() {
	input1 := []int{1, 1, 1, 1}

	fmt.Println(numIdenticalPairs(input1))

}

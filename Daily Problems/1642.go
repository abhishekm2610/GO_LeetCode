package main

import (
	"fmt"
)

func furthestBuilding(heights []int, bricks int, ladders int) int {
	tbricks := bricks
	tladders := ladders
	i := len(heights) - 1
	cur := i

	// possbileans := []int{}
	for i >= 1 {
		if heights[i] > heights[i-1] {
			if heights[i]-heights[i-1] < bricks && bricks != 0 {
				bricks -= heights[i] - heights[i-1]
				i--
			} else if ladders != 0 {
				ladders -= 1
				i--
			} else if i != 0 {
				cur -= 1
				i = cur
				bricks = tbricks
				ladders = tladders
			}
			if i == 0 {
				break
			}
		} else {
			i--
		}
	}
	// for i = 1; i < len(heights); i++ {
	// 	if heights[i] > heights[i-1] {
	// 		if heights[i]-heights[i-1] <= bricks && bricks != 0 {
	// 			bricks -= heights[i] - heights[i-1]
	// 			continue
	// 		} else if ladders != 0 {
	// 			ladders -= 1
	// 			continue
	// 		} else {
	// 			break
	// 		}
	// 	}
	// }
	// j := 0
	// ladders = tladders
	// bricks = tbricks
	// for j = 1; j < len(heights); j++ {

	// 	if heights[j] > heights[j-1] {

	// 		if ladders != 0 {
	// 			ladders -= 1
	// 			continue
	// 		} else if heights[j]-heights[j-1] <= bricks && bricks != 0 {
	// 			bricks -= heights[j] - heights[j-1]
	// 			continue
	// 		} else {
	// 			break
	// 		}
	// 	}
	// }
	return cur
}
func main() {

	ip1 := []int{1, 5, 1, 2, 3, 4, 10000}
	ip2 := 4
	ip3 := 1

	fmt.Println(furthestBuilding(ip1, ip2, ip3))
}

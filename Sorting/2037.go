package main

import (
	"fmt"
	"math"
	"sort"
)

//Original Solution
func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	moves := 0
	for i, v := range students {
		if v == seats[i] {
			continue
		} else {
			moves += int(math.Abs(float64(v) - float64(seats[i])))

		}
	}
	return moves
}

func main() {
	input1 := []int{2, 2, 6, 6}
	intput2 := []int{1, 3, 2, 6}

	fmt.Println(minMovesToSeat(input1, intput2))

}

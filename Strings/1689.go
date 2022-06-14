package main

import (
	"fmt"
	"math"
	"strconv"
)

//Original Solution
func minPartitions(n string) int {
	ans := math.MinInt
	for _, v := range n {
		strInt, _ := strconv.Atoi(string(v))
		if strInt > ans {
			ans = strInt
		}
	}
	return ans
}

func main() {
	input := "82734"

	fmt.Println(minPartitions((input)))
}

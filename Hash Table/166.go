package main

import (
	"fmt"
	"strings"
)

//Original Solution
func fractionToDecimal(numerator int, denominator int) string {
	if denominator == 0 {
		return ("0")
	}
	result := float64(float64(numerator) / float64(denominator))
	strResFrac := fmt.Sprintf("%v", result)
	strResRational := strings.Split(strResFrac, ".")
	strRes := strResRational[1]
	memo := make(map[string]int)

	repeatingIndex := 0
	i := 1
	for i <= len(strRes) {
		if count, isPresent := memo[strRes[repeatingIndex:i]]; isPresent {
			count++
			if count == 2 {
				break
			}
			repeatingIndex = i

			continue

		} else {
			memo[strRes[repeatingIndex:i]] = 1
		}
		i++
	}
	fmt.Println(memo)
	return strResRational[0] + "." + strRes[:repeatingIndex] + "(" + strRes[repeatingIndex:i] + ")"
}
func main() {
	// input1 := 23424242424242
	// intput2 := 4
	input1 := 4
	intput2 := 333
	fmt.Println(fractionToDecimal(input1, intput2))

}

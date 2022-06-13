package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

//Original Solution
func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	} else if len(s) == 1 {
		if s[0] != '-' && s[0] != '+' {

			if '0' <= s[0] && s[0] <= '9' {
				intVar, _ := strconv.ParseInt(string(s[0]), 0, 32)
				return int(intVar)
			}
		} else {
			return 0
		}
	}
	if s[0] != '-' && s[0] != '+' {

		if s[0] > '9' || s[0] < '0' {
			return 0
		}
	}
	numStr := "0"
	sign := 1
	for i, c := range s {

		if c == '-' {
			if '0' <= s[i+1] && s[i+1] <= '9' {
				startingIndex := i + 1
				sign = -1
				for startingIndex < len(s) && '0' <= s[startingIndex] && s[startingIndex] <= '9' {
					numStr += string(s[startingIndex])
					startingIndex++
				}
				break
			} else {
				numStr += string(s[i])
				break
			}

		} else if c == '+' {
			if '0' <= s[i+1] && s[i+1] <= '9' {
				startingIndex := i + 1
				sign = 1
				for startingIndex < len(s) && '0' <= s[startingIndex] && s[startingIndex] <= '9' {
					numStr += string(s[startingIndex])
					startingIndex++
				}
				break
			} else {
				numStr += string(s[i])
				break
			}
		} else if '0' <= c && c <= '9' {
			if '0' <= s[i+1] && s[i+1] <= '9' {
				numStr += string(s[i])
				startingIndex := i + 1
				sign = 1
				for startingIndex < len(s) && '0' <= s[startingIndex] && s[startingIndex] <= '9' {
					numStr += string(s[startingIndex])
					startingIndex++
				}
				break
			} else {
				numStr += string(s[i])
				break
			}
		}
	}
	numStr = strings.TrimLeft(numStr, "0")
	if len(numStr) == 0 {
		return 0
	}
	firstDigit, _ := strconv.ParseInt(string(numStr[0]), 0, 32)
	intValue := sign * int(firstDigit)
	for i := 1; i < len(numStr); i++ {
		intValue *= 10
		intVar, _ := strconv.ParseInt(string(numStr[i]), 0, 32)
		intValue += int(intVar) * sign
		if intValue < math.MaxInt32 {
			continue
		}
		if math.MaxInt32/10 < intValue || (math.MaxInt32/10 == intValue && math.MaxInt32%10 < intVar) {
			if sign == 1 {
				return math.MaxInt32
			}

			return math.MinInt32

		}
	}
	if intValue < int(math.Pow(-2, 31)) {
		intValue = int(math.Pow(-2, 31))
	} else if intValue > int(math.Pow(2, 31))-1 {
		intValue = int(math.Pow(2, 31)) - 1
	}
	return intValue
}

func main() {
	input :=
		"2147483646"
	fmt.Println(myAtoi((input)))
}

package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	if len(numbers) == 2 {
		return []int{1, 2}
	}
	// Incorrect solution
	// i := 0
	// j := len(numbers) - 1
	// for i < j {
	// 	if numbers[j] < target {
	// 		if numbers[j]+numbers[i] == target {
	// 			break
	// 		} else if numbers[j]+numbers[i] < target {
	// 			i++
	// 			continue
	// 		} else {
	// 			j--
	// 		}
	// 	} else if numbers[j] == target && numbers[i] == 0 {
	// 		break

	// 	} else {
	// 		j--
	// 	}
	// }

	i := 0
	j := len(numbers) - 1
	for i < j {
		if numbers[i]+numbers[j] == target {
			break
		} else if numbers[i]+numbers[j] > target {
			j--
			continue
		} else if numbers[i]+numbers[j] < target {
			i++
			continue
		}
	}
	return []int{i + 1, j + 1}
}

func main() {
	ip1 := []int{0, 0, 4, 90}
	ip2 := 0

	fmt.Println(twoSum(ip1, ip2))
}

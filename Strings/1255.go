package main

import (
	"fmt"
	"math"
)

func maxScoreWords(words []string, letters []byte, score []int) int {
	if len(letters) == 0 || len(words) == 0 {
		return 0
	}
	lettersMap := make(map[byte]int)
	for _, v := range letters {
		if _, ok := lettersMap[v]; ok {
			lettersMap[v]++
		} else {
			lettersMap[v] = 0
		}
	}
	ans := math.MinInt
	for _, ov := range words {
		wordsValueMap := make(map[string]int)
		tempLettersMap := make(map[byte]int)

		for key, val := range lettersMap {
			tempLettersMap[key] = val
		}
		for i, _ := range ov {
			if _, ok := tempLettersMap[ov[i]]; ok {
				tempLettersMap[ov[i]] -= 1
			}
		}
		for _, v := range words {
			wordsMap := make(map[byte]int)
			// fmt.Println(ov, v, tempLettersMap)
			// fmt.Println(wordsMap)
			if v != ov {
				for i, _ := range v {
					if _, ok := tempLettersMap[v[i]]; ok {
						wordsMap[v[i]] += 1
						tempLettersMap[v[i]] -= 1
					} else {
						wordsMap = make(map[byte]int)
						break
					}
				}
				fmt.Println(wordsMap)

				if len(wordsMap) == 0 {
					continue
				} else {

					for k, val := range wordsMap {
						if val > tempLettersMap[k] {

							break
						}
					}

					for k, _ := range wordsMap {
						// fmt.Println(wordsValueMap[string(v)], lettersMap[byte(wordsMap[k]+97)])
						wordsValueMap[string(v)] += score[k-97] * wordsMap[k]
					}

				}
			}
		}
		// fmt.Println(wordsValueMap)
		sum := 0
		for _, v := range wordsValueMap {
			sum += v
		}
		if sum > ans {
			ans = sum
		}
	}

	return ans
}

func main() {
	words := []string{"dog", "cat", "dad", "good"}
	letters := []byte{'a', 'a', 'c', 'd', 'd', 'd', 'g', 'o', 'o'}
	score := []int{1, 0, 9, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	fmt.Println(maxScoreWords(words, letters, score))
}

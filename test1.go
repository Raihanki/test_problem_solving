package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{1, 3, 5, 7, 9}
	fmt.Println(MiniMaxSum(numbers))
}

func MiniMaxSum(numbers []int) (int, int) {
	var results []int

	for i := 0; i < len(numbers); i++ {
		res := 0
		for j := 0; j < len(numbers); j++ {
			if i != j {
				res += numbers[j]
			}
		}
		results = append(results, res)
		res = 0
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i] < results[j]
	})

	return results[0], results[len(results)-1]
}

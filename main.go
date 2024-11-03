package main

import (
	"fmt"
)

func findDuplicates(nums []int) []int {
	duplicateMap := make(map[int]int)
	duplicates := []int{}
	for _, num := range nums {
		if _, exists := duplicateMap[num]; exists {
			duplicates = append(duplicates, num)
		} else {
			duplicateMap[num] = 1
		}
	}
	return duplicates
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	duplicates := findDuplicates(nums)
	fmt.Println("Duplicate numbers:", duplicates)
}

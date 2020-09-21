package main

import (
	"fmt"
)

func twoSum1(nums []int, target int) []int {
	numLen := len(nums)

	if numLen > 1 {
		for i := 0; i < numLen-1; i++ {
			for j := i + 1; j < numLen-1; j++ {
				if nums[i]+nums[j] == target {
					return []int{i, j}
				}
			}
		}
	}
	return []int{0, 0}
}

func twoSum(nums []int, target int) []int {
	m := map[int]int{}

	numLen := len(nums)
	if numLen > 1 {
		for index, num := range nums {
			if i, ok := m[target-num]; ok {
				return []int{i, index}
			}
			m[num] = index
		}
	}
	return []int{0, 0}
}

func main() {
	nums := []int{-10, -1, -18, -19}
	target := -10
	fmt.Print(twoSum(nums, target))
}

func main() {
	nums := []int{-1, 2, -3, -4, 5}
	target := -8
	fmt.Println(twoSum(nums, target))
}

package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var result [][]int

	if len(nums) < 3 {
		return result
	}
	sort.Ints(nums)
	fmt.Println(nums)

	i := 0

	for i < len(nums) {

		if i > 0 {
			for i < len(nums) && (nums[i] == nums[i-1]) {
				i++
			}
		}
		fmt.Println(i)

		val := nums[i]
		left, right := i+1, len(nums)-1
		for left < right && left < len(nums) {

			if nums[left]+nums[right]+val == 0 {
				fmt.Println("before", i, left, right, " :", val, nums[left], nums[right])
				result = append(result, []int{val, nums[left], nums[right]})
				for left < len(nums)-1 && (nums[left] == nums[left+1]) {
					left++
				}
				left++
				for right > 0 && (nums[right] == nums[right-1]) {
					right--
				}
				right--
				fmt.Println(i, left, right)
			} else if nums[left]+nums[right]+val < 0 {
				for left < len(nums)-1 && (nums[left] == nums[left+1]) {
					left++
				}
				left++
			} else {
				for right > 0 && (nums[right] == nums[right-1]) {
					right--
				}
				right--
			}
		}
		i++
	}

	return result
}

func main() {
	fmt.Println(threeSum([]int{-4, -3, -2, -1, -1, 0, 0, 1, 2, 3, 4}))
}

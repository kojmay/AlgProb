package main

import (
	"fmt"
)

func maxArea(height []int) int {
	maxArea := 0
	start, end := 0, len(height)-1

	for end > start {
		area := (end - start) * min(height[end], height[start])
		maxArea = max(maxArea, area)
		if height[end] > height[start] {
			start++
		} else {
			end--
		}
	}

	return maxArea
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	var height []int
	height = []int{1, 2, 1}
	fmt.Print(maxArea(height))
}

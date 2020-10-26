package main

import (
	"fmt"
)

func maximalNetworkRank(n int, roads [][]int) int {
	// rows, cols := len(roads), len(roads[0])
	connected := make([][]bool, n)
	for i := range connected {
		connected[i] = make([]bool, n)
	}
	result := 0
	conNums := make([]int, n)

	for _, item := range roads {
		conNums[item[0]]++
		conNums[item[1]]++
		connected[item[0]][item[1]] = true
		connected[item[1]][item[0]] = true
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if connected[i][j] {
				result = max(result, conNums[i]+conNums[j]-1)
			} else {
				result = max(result, conNums[i]+conNums[j])
			}
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	n := 4
	roads := [][]int{{0, 1}, {0, 3}, {1, 2}, {1, 3}}

	fmt.Print(maximalNetworkRank(n, roads))
}

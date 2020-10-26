package main

import (
	"fmt"
)

// solution1 : use DP
func longestValidParentheses(s string) int {
	result := 0
	if len(s) <= 1 {
		return result
	}

	var dp []int
	for i := 0; i < len(s); i++ {
		if i == 0 {
			dp = append(dp, 0)
		} else {
			// fmt.Println(string(s[i]), i)
			if s[i] == ')' && s[i-1] == '(' {
				if i >= 2 {
					dp = append(dp, dp[i-2]+2)
				} else {
					dp = append(dp, 2)
				}
			} else if s[i] == ')' && s[i-1] == ')' {
				// fmt.Println(i-dp[i-1]-1, string(s[i-dp[i-1]-1]))
				if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
					if i-dp[i-1] >= 2 {
						// fmt.Println(i - dp[i-1] - 2)
						dp = append(dp, dp[i-1]+dp[i-dp[i-1]-2]+2)
					} else {
						dp = append(dp, dp[i-1]+2)
					}
				} else {
					dp = append(dp, 0)
				}
			} else if s[i] == '(' {
				dp = append(dp, 0)
			}
		}
		if dp[i] > result {
			result = dp[i]
		}
	}

	return result
}

func main() {

	fmt.Println(longestValidParentheses(")()())"))
}

package main

import "fmt"

func generateParenthesis(n int) []string {
	result := []string{}

	subGenerateParenthesis(n, 0, 0, "", &result)

	return result
}

func subGenerateParenthesis(n int, left int, right int, str string, result *[]string) {
	if left == n && right == n {
		*result = append(*result, str)
		return
	}

	if left < right || left > n {
		return
	}

	subGenerateParenthesis(n, left+1, right, str+"(", result)
	subGenerateParenthesis(n, left, right+1, str+")", result)
}

func main() {
	fmt.Println(generateParenthesis(3))
}

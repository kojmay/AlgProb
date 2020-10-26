package main

import "fmt"

func maxDepth(s string) int {
	if len(s) <= 1 {
		return 0
	}
	result := 0
	var stack []string

	for _, val := range s {
		item := string(val)
		if item == "(" {
			stack = append(stack, "(")
			if result < len(stack) {
				result = len(stack)
			}
			// fmt.Println(stack, result)
		} else if item == ")" {
			stack = stack[:len(stack)-1]
			// fmt.Println(stack, result)
		}
	}

	return result
}

func main() {
	fmt.Println(maxDepth("(1+(2*3)+((8)/4))+1"))
}

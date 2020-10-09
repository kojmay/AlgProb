package main

import "fmt"

func minOperations(logs []string) int {
	depth := 0
	for _, v := range logs {
		switch v {
		case "../":
			if depth > 0 {
				depth--
			}
		case "./":
		default:
			depth++
		}
	}
	return depth
}

func main() {
	logs := []string{"d1/", "../", "../", "../"}
	fmt.Println(minOperations(logs))

	logs = []string{"d1/", "d2/", "./", "d3/", "../", "d31/"}
	fmt.Println(minOperations(logs))

	logs = []string{"d1/", "d2/", "../", "d21/", "./"}
	fmt.Println(minOperations(logs))

}

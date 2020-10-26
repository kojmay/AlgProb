package main

import "fmt"

func checkPalindromeFormation(a string, b string) bool {

	return subCheckPalindromeFormation(a, b) || subCheckPalindromeFormation(b, a)
}

func subCheckPalindromeFormation(a string, b string) bool {
	if subCheckSinglePalindromeFormation(a) || subCheckSinglePalindromeFormation(b) {
		return true
	}

	lenA, lenB := len(a), len(b)
	startA, startB := 0, lenB-1

	for startA < lenA && startB >= 0 {
		if a[startA] == b[startB] {
			if startA+1 == startB || startA+2 == startB {
				return true
			}
			startA++
			startB--
		} else {
			subA := a[startA : startB+1]
			subB := b[startA : startB+1]

			return subCheckSinglePalindromeFormation(subA) || subCheckSinglePalindromeFormation(subB)
		}
	}

	return false
}

func subCheckSinglePalindromeFormation(a string) bool {
	endA := len(a) - 1
	startA := 0
	for startA < endA {
		if a[startA] != a[endA] {
			return false
		}
		startA++
		endA--
	}
	return true
}

func main() {

	fmt.Println(checkPalindromeFormation("cdeoo", "oooab"))
	fmt.Println(checkPalindromeFormation("aejbaalflrmkswrydwdkdwdyrwskmrlfqizjezd", "uvebspqckawkhbrtlqwblfwzfptanhiglaabjea"))

	//"aejbaalflrmkswrydwdkdwdyrwskmrlfqizjezd"
	//"uvebspqckawkhbrtlqwblfwzfptanhiglaabjea"
}

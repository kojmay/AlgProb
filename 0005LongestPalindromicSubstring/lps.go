package main

import "fmt"

func longestPalindrome(s string) string {
	// subLongestPalindrome(s, 0, &start, &end, &length)
	index, length := 0, 0
	for i := 0; i < len(s); i++ {
		start1, length1 := subLongestPalindrome(s, i-1, i+1)
		start2, length2 := subLongestPalindrome(s, i, i+1)
		if length1 > length {
			length = length1
			index = start1
		}
		if length2 > length {
			length = length2
			index = start2
		}
	}

	return s[index : index+length]
}

func subLongestPalindrome(s string, start int, end int) (int, int) {
	i, j := start, end
	for i >= 0 && j < len(s) {
		if s[i] != s[j] {
			return i + 1, j - i - 1
		}
		i--
		j++
	}
	return i + 1, j - i - 1
}

func main() {
	fmt.Println(longestPalindrome("cbbd"))
}

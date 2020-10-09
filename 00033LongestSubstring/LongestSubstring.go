package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	result := 0
	dict := make(map[rune]int)
	start := 0

	// for i := 0; i < len(s); i++ {
	// 	if v, ok := dict[s[i:i+1]]; ok && v > start {
	// 		start = v + 1
	// 	}
	// 	dict[s[i:i+1]] = i

	// 	if result < i-start+1 {
	// 		result = i - start + 1
	// 	}
	// 	// fmt.Println(result, dict, start, end)
	// }

	for idx, c := range s {
		if v, ok := dict[c]; ok && v >= start {
			start = v + 1
		}
		dict[c] = idx

		if result < idx-start+1 {
			result = idx - start + 1
		}
		// fmt.Println(start, idx, dict, result)
	}

	return result
}

func main1() {

	fmt.Println(lengthOfLongestSubstring("aabaab!bb"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring(""))

}

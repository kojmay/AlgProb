package main

import (
	"fmt"
	"strconv"
)

func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}

	// fmt.Println()
	s := strconv.Itoa(x)
	for left, right := 0, len(s)-1; left < right; {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func isPalindrome(x int) bool {
	if x < 0 || x%10 == 0 { //小于0 或者以0结尾
		return false
	}

	y := 0
	// var y int
	for ; x > y; x /= 10 {
		y = x%10 + y*10
	}

	return x == y || x == y/10
}

func main2() {

	fmt.Println(isPalindrome(121))

	// fmt.Println(strconv.Itoa(121))
	// fmt.Println(strconv.Itoa(-1))
}

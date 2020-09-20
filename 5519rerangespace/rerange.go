package main

import (
	"fmt"
	"strings"
)

var p = fmt.Println

func reorderSpaces(text string) string {

	spaceNum := strings.Count(text, " ")
	if spaceNum == 0 {
		return text
	}

	words := []string{}

	for _, s := range strings.Split(text, " ") {
		if !strings.Contains(s, " ") {
			words = append(words, s)
		}
	}

	wordNum := len(words)

	if wordNum == 1 {
		return words[0] + strings.Repeat(" ", spaceNum)
	}

	eachNum := spaceNum / (wordNum - 1)
	appendNum := spaceNum % (wordNum - 1)
	result := ""
	for i, word := range words {
		result = result + word
		if i == wordNum-1 {
			result += strings.Repeat(" ", appendNum)
		} else {
			result += strings.Repeat(" ", eachNum)
		}
	}
	return result
}

func main() {

	fmt.Println(reorderSpaces("  this   is  a sentence "))
}

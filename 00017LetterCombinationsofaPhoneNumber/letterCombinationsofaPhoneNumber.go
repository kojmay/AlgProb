package main

import "fmt"

func letterCombinations(digits string) []string {
	var result []string
	if len(digits) == 0 {
		return result
	}

	str := ""
	subLetterCombinations(digits, 0, str, &result)

	return result
}

func subLetterCombinations(digits string, index int, str string, result *[]string) {
	if index == len(digits) {
		// item := *str
		*result = append(*result, str)
		// fmt.Println(*result)
		return
	} else {
		digitLetters := digitLetter(string(digits[index]))
		// fmt.Println(string(digits[index]), digitLetters)
		if len(digitLetters) == 0 {
			subLetterCombinations(digits, index+1, str, result)
			return
		}
		for _, val := range digitLetters {
			str = str + val
			subLetterCombinations(digits, index+1, str, result)
			str = str[:len(str)-1]
		}
	}
}

func digitLetter(a string) []string {
	switch a {
	case "2":
		return []string{"a", "b", "c"}
	case "3":
		return []string{"d", "e", "f"}
	case "4":
		return []string{"g", "h", "i"}
	case "5":
		return []string{"j", "k", "l"}
	case "6":
		return []string{"m", "n", "o"}
	case "7":
		return []string{"p", "q", "r", "s"}
	case "8":
		return []string{"t", "u", "v"}
	case "9":
		return []string{"w", "x", "y", "z"}
	default:
		return []string{}
	}
}

func main() {
	fmt.Print(letterCombinations("123"))
}

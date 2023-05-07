package src

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	splitted := strings.Split(s, " ")

	for i := 0; i < len(splitted); i++ {
		word := splitted[i]
		left := 0
		right := len(word) - 1
		wordChars := strings.Split(word, "")

		for left < right {
			wordChars[left], wordChars[right] = wordChars[right], wordChars[left]
			left++
			right--
		}

		splitted[i] = strings.Join(wordChars, "")
	}

	return strings.Join(splitted, " ")
}

func rreverseWords_solve() {
	variable := "Let's take LeetCode contest"
	solution := reverseWords(variable)
	fmt.Println("The the result is: ", solution)
}

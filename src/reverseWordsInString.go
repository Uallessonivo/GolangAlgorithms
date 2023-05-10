package src

import (
	"fmt"
	"strings"
)

/*
Reverse Words in a String III

Given a string s, reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.


Example 1:
Input: s = "Let's take LeetCode contest"
Output: "s'teL ekat edoCteeL tsetnoc"

Example 2:
Input: s = "God Ding"
Output: "doG gniD"
*/

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

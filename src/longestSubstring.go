package src

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	start := 0
	longest := 0
	set := make(map[rune]bool)

	for i, c := range s {
		if set[c] {
			for s[start] != byte(c) {
				delete(set, rune(s[start]))
				start++
			}
			start++
		}

		set[c] = true

		if i-start+1 > longest {
			longest = i - start + 1
		}
	}

	return longest
}

func lengthOfLongestSubstring_solve() {
	variable := "abcabcbb"
	solution := lengthOfLongestSubstring(variable)
	fmt.Println("The the result is: ", solution)
}

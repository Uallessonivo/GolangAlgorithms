package src

import (
	"fmt"
)

/*
Longest Substring Without Repeating Characters
Given a string s, find the length of the longest substring without repeating characters.


Example 1:
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.


Example 2:
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.


Example 3:
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

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

func LengthOfLongestSubstring_solve() {
	variable := "abcabcbb"
	solution := lengthOfLongestSubstring(variable)
	fmt.Println("The the result is: ", solution)
}

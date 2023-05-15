package src

import (
	"fmt"
)

/*
Permutation in String

Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.
In other words, return true if one of s1's permutations is the substring of s2.


Example 1:
Input: s1 = "ab", s2 = "eidbaooo"
Output: true
Explanation: s2 contains one permutation of s1 ("ba").

Example 2:
Input: s1 = "ab", s2 = "eidboaoo"
Output: false
*/

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	freq := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		freq[s1[i]]++
	}

	left, right := 0, 0

	for right < len(s2) {
		if freq[s2[right]] > 0 {
			freq[s2[right]]--
			right++
		} else {
			freq[s2[left]]++
			left++
		}

		if right-left == len(s1) {
			return true
		}
	}

	return false
}

func CheckInclusion_solve() {
	solution := checkInclusion("ab", "eidbaooo")
	fmt.Println("The the result is: ", solution)
}

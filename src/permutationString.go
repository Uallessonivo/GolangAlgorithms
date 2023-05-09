package src

import (
	"fmt"
)

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

func checkInclusion_solve() {
	solution := checkInclusion("ab", "eidbaooo")
	fmt.Println("The the result is: ", solution)
}

package practice

/*
Valid Anagram

Given two strings s and t, return true if t is an anagram of s, and false otherwise.
An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.


Example 1:
Input: s = "anagram", t = "nagaram"
Output: true

Example 2:
Input: s = "rat", t = "car"
Output: false
*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	set := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		set[s[i]]++
		set[t[i]]--
	}

	for _, v := range set {
		if v != 0 {
			return false
		}
	}

	return true
}

package practice

/*
Find the Difference

You are given two strings s and t.

String t is generated by random shuffling string s and then add one more letter at a random position.

Return the letter that was added to t.



Example 1:
Input: s = "abcd", t = "abcde"
Output: "e"

Explanation: 'e' is the letter that was added.

Example 2:
Input: s = "", t = "y"
Output: "y"
*/

func FindTheDifference(s string, t string) byte {
	result := byte(0)

	for i := 0; i < len(s); i++ {
		result = result ^ s[i]
	}

	for i := 0; i < len(t); i++ {
		result = result ^ t[i]
	}

	return result
}

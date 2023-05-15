package practice

/*
Find the Index of the First Occurrence in a String

Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:
Input: haystack = "sadbutsad", needle = "sad"
Output: 0

Explanation: "sad" occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.

Example 2:
Input: haystack = "leetcode", needle = "leeto"
Output: -1

Explanation: "leeto" did not occur in "leetcode", so we return -1.
*/

func StrStr(haystack string, needle string) int {
	// Declare m and n as variables, and initialize them with the length of needle and haystack respectively.
	m := len(needle)
	n := len(haystack)

	// Declare the windowStart variable, and initialize it with 0.
	// Iterate windowStart till starting index of the last substring of length m, i.e till n - m.
	for windowStart := 0; windowStart <= n-m; windowStart++ {

		// For each windowStart, iterate variable i from 0 to m - 1
		for i := 0; i < m; i++ {

			// Check if the character in the window i.e index windowStart + i is equal to the the character in the needle.
			if needle[i] == haystack[windowStart+i] {
				// Since they are equal, we break out of the loop and move on to the next iteration of the outer loop.
				break
			}

			// If the loop over i completes without any mismatches, then the function returns windowStart.
			// Which is the index of the first occurrence of needle in haystack.
			if i == m-1 {
				return windowStart
			}
		}
	}

	// If the outer loop over windowStart completes without finding any matches.
	// The function returns -1, indicating that needle does not occur in haystack.
	return -1
}

package src

import "fmt"

func reverseString(s []byte) []byte {
	left := 0
	right := len(s) - 1

	for left <= right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}

	return s
}

func reverseString_solve() {
	variable := []byte{'h', 'e', 'l', 'l', 'o'}
	solution := reverseString(variable)
	fmt.Println("The the result is: ", solution)
}

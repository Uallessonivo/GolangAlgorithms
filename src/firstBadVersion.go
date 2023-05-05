package src

import "fmt"

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(version int) bool {
	if version < 3 {
		return true
	}

	return false
}

func firstBadVersion(n int) int {
	left := 0
	right := n

	for left < right {
		mid := (left + right) / 2

		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func firstBadVersion_solve() {
	variable := 5
	solution := firstBadVersion(variable)
	fmt.Println("The the result is: ", solution)
}

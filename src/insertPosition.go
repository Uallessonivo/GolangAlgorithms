package src

import "fmt"

func insertPosition(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

func insertPosition_solve() {
	variable := []int{1, 3, 5, 6}
	solution := insertPosition(variable, 7)
	fmt.Println("The the result is: ", solution)
}

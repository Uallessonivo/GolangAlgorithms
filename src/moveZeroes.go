package src

import "fmt"

/*
Move Zeroes

Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
Note that you must do this in-place without making a copy of the array.


Example 1:
Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]

Example 2:
Input: nums = [0]
Output: [0]
*/

func moveZeroes(nums []int) []int {
	n := len(nums)
	index := 0

	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[i], nums[index] = nums[index], nums[i]
			index++
		}
	}

	return nums
}

func moveZeroes_solve() {
	variable := []int{0, 1, 0, 3, 12}
	solution := moveZeroes(variable)
	fmt.Println("The the result is: ", solution)
}

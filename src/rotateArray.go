package src

import "fmt"

/*
Rotate Array

Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.


Example 1:
Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]

Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]

Example 2:
Input: nums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]

Explanation:
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]
*/

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func rotate(nums []int, k int) []int {
	n := len(nums)
	k %= n

	reverse(nums, 0, n-1) // Inverte o array inteiro
	reverse(nums, 0, k-1) // Inverte os primeiros k elementos
	reverse(nums, k, n-1) // Inverte os elementos restantes

	return nums
}

func Rotate_solve() {
	variable := []int{1, 2, 3, 4, 5, 6, 7}
	solution := rotate(variable, 3)
	fmt.Println("The the result is: ", solution)
}

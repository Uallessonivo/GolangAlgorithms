package src

import "fmt"

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

func rotate_solve() {
	variable := []int{1, 2, 3, 4, 5, 6, 7}
	solution := rotate(variable, 3)
	fmt.Println("The the result is: ", solution)
}

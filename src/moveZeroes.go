package src

import "fmt"

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

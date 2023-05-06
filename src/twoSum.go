package src

import "fmt"

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		} else if numbers[left]+numbers[right] < target {
			left++
		} else {
			right--
		}
	}

	return nil
}

func twoSum_solve() {
	variable := []int{2, 7, 11, 15}
	solution := twoSum(variable, 9)
	fmt.Println("The the result is: ", solution)
}

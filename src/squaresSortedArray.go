package src

import (
	"fmt"
	"sort"
)

func sortedSquares(nums []int) []int {
	squares := []int{}

	for _, num := range nums {
		squares = append(squares, num*num)
	}

	sort.Ints(squares)

	return squares
}

func sortedSquares_solve() {
	variable := []int{-4, -1, 0, 3, 10}
	solution := sortedSquares(variable)
	fmt.Println("The the result is: ", solution)
}

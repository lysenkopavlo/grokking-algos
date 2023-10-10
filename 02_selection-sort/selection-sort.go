package main

import "fmt"

// findMinIndex finds the smallest value in the given slice
// and returns its index
func findMinIndex(nums []int) int {
	min := nums[0]
	minIndex := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
			minIndex = i
		}
	}
	return minIndex
}

// selectionSort returns sorted slice
func selectionSort(nums []int) []int {
	// create resulting slice
	sorted := make([]int, len(nums))

	for i := range nums {
		minIndex := findMinIndex(nums[i:]) // decrease numbers of elements every iteration

		// reason add i to minIndex
		// minIndex is behind i
		sorted[i] = nums[minIndex+i]

		// swap elements to exclude min value
		// from findMinIndex
		nums[i], nums[minIndex+i] = nums[minIndex+i], nums[i]
	}

	return sorted
}

func main() {
	fmt.Println(selectionSort([]int{20, 50, 30, 10, 60}))
}

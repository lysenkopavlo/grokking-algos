package main

import "fmt"

// binarySearch works with sorted slice
// and returns an index of element,
// if it's in slice
func binarySearch(nums []int, item int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == item {
			return mid
		}
		if nums[mid] > item {
			high = mid - 1
		}
		if nums[mid] < item {
			low = mid + 1
		}
	}

	return -1
}

func main() {
	fmt.Println(binarySearch([]int{1, 3, 5, 7, 9, 10}, 7))
	fmt.Println(binarySearch([]int{1, 3, 5, 7, 9, 10}, 0))
}

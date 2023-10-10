package main

import "fmt"

// first exercise:
// sum returns a sum of elements in slice
func sum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return nums[0] + sum(nums[1:])
}

// findMax recursively finds maximum value
// and returns it
func findMax(nums []int) int {
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}

	subMax := findMax(nums[1:])

	if nums[0] > subMax {
		return nums[0]
	}
	return subMax
}

// quickSort returns sorted slice from min to max
func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	} else {
		pivot := nums[0]

		left := make([]int, 0, len(nums)/2)
		right := make([]int, 0, len(nums)/2)

		for _, num := range nums[1:] {
			if pivot > num {
				left = append(left, num)
			} else {
				right = append(right, num)
			}
		}
		left = append(quickSort(left), pivot)
		right = quickSort(right)
		return append(left, right...)
	}

}

func main() {
	s := []int{40, 30, 10, 20} // sum ==100

	fmt.Println(sum(s))
	fmt.Println(findMax(s))
	fmt.Println(quickSort(s))

}

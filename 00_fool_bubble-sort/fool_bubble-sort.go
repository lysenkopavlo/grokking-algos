package main

import (
	"fmt"
	"slices"
)

func foolSort(nums []int) {
	for i := 0; i < len(nums)-1; {
		if nums[i] > nums[i+1] {
			nums[i], nums[i+1] = nums[i+1], nums[i]
			i = 0
		} else {
			i++
		}
	}
}

func bubbleSort(nums []int) {
	flag := false
	for !flag {
		flag = true
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				flag = false
			}
		}
	}
}

func main() {
	nums := []int{50, 40, 30, 20, 10}
	foolSort(nums)
	fmt.Println(nums, "sorted:", slices.IsSorted(nums))

	nums = []int{51, 41, 31, 21, 11}
	bubbleSort(nums)
	fmt.Println(nums, "sorted: ", slices.IsSorted(nums))

}

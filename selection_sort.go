package main

import (
	"fmt"
)

func selectionSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		minimum := i
		for j := i; j < len(nums); j++ {
			if nums[j] < nums[minimum] {
				minimum = j
			}
		}
		nums[i], nums[minimum] = nums[minimum], nums[i] // swap
	}
	return nums
}

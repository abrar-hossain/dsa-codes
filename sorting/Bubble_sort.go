
package main

import "fmt"

func bubbleSort(nums []int) []int {
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= i-1; j++ {
			if nums[j] > nums[j+1] {
				// swap nums[j] and nums[j+1]
				temp := nums[j+1]
				nums[j+1] = nums[j]
				nums[j] = temp
			}
		}
	}
	return nums
}

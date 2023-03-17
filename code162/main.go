package main

import "fmt"

func findPeakElement(nums []int) int {
	i, j := 0, len(nums)-1
	for i < j {
		m := (i + j) >> 1
		if m+1 < len(nums) {
			if nums[m] <= nums[m+1] {
				i = m + 1
			} else {
				j = m
			}
		} else {
			break
		}
	}
	return i
}

func main() {
	ret1 := findPeakElement([]int{1, 2, 3, 1})
	if ret1 != 2 {
		panic(fmt.Errorf("code162 error, ret1 != 2: %d", ret1))
	}
}

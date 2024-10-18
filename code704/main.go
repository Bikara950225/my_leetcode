package main

import "fmt"

func search(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i < j {
		m := (i + j) >> 1
		if target > nums[m] {
			i = m + 1
		} else {
			j = m
		}
	}
	if nums[i] == target {
		return i
	}
	return -1
}

func main() {
	targetIndex := search([]int{1, 2, 3, 4, 5}, 5)
	if targetIndex != 4 {
		panic(fmt.Errorf("code704 error, not find targetIndex: %d, expect: %d",
			targetIndex, 4,
		))
	}

	targetIndex = search([]int{1, 2, 3, 5, 6}, 7)
	if targetIndex != -1 {
		panic(fmt.Errorf("code704 error, not find targetIndex: %d, expect: %d",
			targetIndex, -1,
		))
	}

	targetIndex = search([]int{1, 2, 3, 5, 6}, 4)
	if targetIndex != -1 {
		panic(fmt.Errorf("code704 error, not find targetIndex: %d, expect: %d",
			targetIndex, -1,
		))
	}
}

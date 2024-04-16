package main

import (
	"fmt"
	"reflect"
)

// nums: [5,7,7,8,8,10], target: 8 => ret: [3, 4]
func searchRange(nums []int, target int) []int {
	i1 := find1(nums, target)
	if nums[i1] != target {
		return []int{-1, -1}
	}

	i2 := find1(nums, target+1)
	if nums[i2] > target {
		i2 -= 1
	}
	return []int{i1, i2}
}

func findMiddle(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i < j {
		m := (i + j) >> 1
		if nums[m] >= target {
			j = m
		} else {
			i = m + 1
		}
	}
	return i
}

func find1(src []int, target int) int {
	i, j := 0, len(src)-1
	for i < j {
		m := (i + j) >> 1
		if src[m] < target {
			i = m + 1
		} else {
			j = m
		}
	}
	return i
}

func main() {
	ret1 := searchRange([]int{5, 7, 7, 8, 8, 10}, 8)
	if !reflect.DeepEqual(ret1, []int{3, 4}) {
		panic(fmt.Errorf("code34 error, not expect ret1: %v", ret1))
	}
}

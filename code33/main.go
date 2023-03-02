package main

import "fmt"

func search(nums []int, target int) int {
	// 思路：旋转后的数组，分为2部分的升序数组，我们需要先找出target属于前半部分还是后半部分的数组，再用二分查找找出目标的index
	b, t := searchSubNums(nums, target)
	return binarySearch(nums, b, t, target)
}

func searchSubNums(nums []int, target int) (b int, t int) {
	b, t = 0, len(nums)-1
	if nums[b] < nums[t] {
		return
	}
	mark := nums[b]
	for b < t {
		m := b + ((t - b) >> 1)
		if nums[m] >= mark {
			b = m + 1
		} else {
			t = m
		}
	}
	if target < mark {
		return b, len(nums) - 1
	} else {
		return 0, b - 1
	}
}

func binarySearch(nums []int, b, t int, target int) int {
	for b < t {
		m := b + ((t - b) >> 1)
		if nums[m] >= target {
			t = m
		} else {
			b = m + 1
		}
	}
	if nums[b] == target {
		return b
	}
	return -1
}

func main() {
	ret1 := search([]int{3, 1}, 1)
	if ret1 != 1 {
		panic(fmt.Errorf("code33 error, 不符合期望值 1: %d", ret1))
	}

	ret2 := search([]int{4, 5, 6, 7, 0, 1, 2}, 0)
	if ret2 != 4 {
		panic(fmt.Errorf("code33 error, 不符合期望值 4: %d", ret2))
	}
}

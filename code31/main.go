package main

import (
	"fmt"
	"reflect"
)

func nextPermutation(nums []int) {
	downI := -1
	for i := len(nums) - 1; i > 0; i-- {
		// 找到递减的下标
		if nums[i] > nums[i-1] {
			downI = i - 1
			break
		}
	}
	// 当前值属于最大值时，downI就是-1
	if downI == -1 {
		reverse(nums, 0, len(nums)-1)
		return
	}

	for i := len(nums) - 1; i > downI; i-- {
		if nums[i] > nums[downI] {
			nums[i], nums[downI] = nums[downI], nums[i]
			break
		}
	}
	reverse(nums, downI+1, len(nums)-1)
}

func reverse(nums []int, begin, end int) {
	for begin < end {
		nums[begin], nums[end] = nums[end], nums[begin]
		begin++
		end--
	}
}

func main() {
	src1 := []int{3, 1, 2}
	nextPermutation(src1)
	expectRet1 := []int{3, 2, 1}
	if !reflect.DeepEqual(src1, expectRet1) {
		panic(fmt.Errorf("code31 error, not expect result: %+v", src1))
	}

	src2 := []int{3, 2, 1}
	nextPermutation(src2)
	expectRet2 := []int{1, 2, 3}
	if !reflect.DeepEqual(src2, expectRet2) {
		panic(fmt.Errorf("code31 error, not expect result: %+v", src2))
	}
}

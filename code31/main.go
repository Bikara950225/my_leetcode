package main

import (
	"fmt"
	"reflect"
)

func nextPermutation(nums []int) {
	downI := -1
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			downI = i - 1
			break
		}
	}
	if downI == -1 {
		reverse(nums)
		return
	}

	for i := len(nums) - 1; i > downI; i-- {
		if nums[i] > nums[downI] {
			nums[i], nums[downI] = nums[downI], nums[i]
			break
		}
	}
	reverse(nums[downI+1:])
}

func reverse(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func main() {
	src1 := []int{3, 1, 2}
	nextPermutation(src1)
	expectRet1 := []int{3, 2, 1}
	if !reflect.DeepEqual(src1, expectRet1) {
		panic(fmt.Errorf("code31 error, not expect result: %+v", src1))
	}
}

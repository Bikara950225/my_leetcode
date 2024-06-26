package main

import (
	"fmt"
	"reflect"
)

// [1,2,1] => [2,-1,2]
// [1,2,3,4,3] => [2,3,4,-1,4]
func nextGreaterElements(nums []int) (ret []int) {
	n := len(nums)

	ret = make([]int, n)
	for i := range ret {
		ret[i] = -1
	}
	// di jian
	var stack []int
	for i := 0; i < 2*n; i++ {
		item := nums[i%n]
		for len(stack) > 0 && item > nums[stack[len(stack)-1]] {
			popIdx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ret[popIdx] = item
		}
		stack = append(stack, i%n)
	}
	return
}

func main() {
	ret1 := nextGreaterElements([]int{1, 2, 1})
	want1 := []int{2, -1, 2}
	if !reflect.DeepEqual(ret1, want1) {
		panic(fmt.Errorf("ret1 error, got != want: %v", want1))
	}
}

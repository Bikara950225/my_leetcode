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
	// 这是一个单调递减的栈，记录nums的下标
	var stack []int
	for i := 0; i < 2*n; i++ {
		for len(stack) > 0 && nums[i%n] > nums[stack[len(stack)-1]] {
			ret[stack[len(stack)-1]] = nums[i%n]
			// pop
			stack = stack[:len(stack)-1]
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

	fmt.Printf("%v\n", nextGreaterElements([]int{1, 3, 2, 1}))
}

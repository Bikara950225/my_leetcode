package main

import (
	"fmt"
	"my_leetcode/internal/tools"
)

func trap(height []int) int {
	// 用DP的话，时间复杂度为 O(n)；空间复杂度为 O(n)
	ret := 0

	// 用于记录下标i左边的最大高度的dp
	leftDp := make([]int, len(height))
	leftDp[0] = height[0]
	for i := 1; i < len(height); i++ {
		leftDp[i] = tools.IntMax(height[i], leftDp[i-1])
	}
	// 用于记录下标i右边的最大高度的dp
	rightDp := make([]int, len(height))
	rightDp[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rightDp[i] = tools.IntMax(height[i], rightDp[i+1])
	}

	for i := 0; i < len(height); i++ {
		// 找出当前下标中，左、右高边的最小的那一个边(木桶原理，取短板)
		minHeight := tools.IntMin(leftDp[i], rightDp[i])
		ret += minHeight - height[i]
	}
	return ret
}

func main() {
	height1 := []int{
		0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1,
	}
	ret1 := trap(height1)
	expectRet1 := 6
	if ret1 != expectRet1 {
		panic(fmt.Errorf("code42 error, not expect result: %d", ret1))
	}
}

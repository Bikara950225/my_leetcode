package main

import (
	"fmt"
	"my_leetcode/internal/tools"
)

func main() {

	ret1 := largestRectangleArea([]int{4, 2, 0, 3, 2, 5})
	if ret1 != 6 {
		panic(fmt.Errorf("code84 error, ret1 != 6: %d", ret1))
	}
}

func largestRectangleArea(heights []int) (ret int) {
	heights = append(heights, 0)
	// 单调递增栈
	var stack []int

	for i, h := range heights {
		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
			topIdx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			var area int
			if len(stack) > 0 {
				area = heights[topIdx] * (i - stack[len(stack)-1] - 1)
			} else {
				area = heights[topIdx] * i
			}
			if area > ret {
				ret = area
			}
		}
		stack = append(stack, i)
	}
	return
}

// 这个方法创建的dp会导致OOM。。。
func dpMethod(heights []int) (ret int) {
	var maxLen int
	for _, height := range heights {
		if height > maxLen {
			maxLen = height
		}
	}
	// 第i个柱的j高度，累计的横长
	dp := make([][]int16, len(heights))
	for i := range dp {
		dp[i] = make([]int16, maxLen+1)
	}

	for i, height := range heights {
		tmp := -1
		for j := 1; j <= height; j++ {
			if i == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + 1
			}
			if tmp == -1 {
				tmp = int(dp[i][j])
			} else {
				tmp = tools.IntMin(tmp, int(dp[i][j]))
			}

			area := tmp * j
			if area > ret {
				ret = area
			}
		}
	}
	return
}

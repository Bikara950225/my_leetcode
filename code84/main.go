package main

import "my_leetcode/internal/tools"

func main() {

}

func largestRectangleArea(heights []int) (ret int) {

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

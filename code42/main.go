package main

// trapWithDP 基于动态规划
// 基于x轴计算每个height的高度差，累加
func trapWithDP(height []int) (ret int) {
	// 用DP的话，时间复杂度为 O(n)；空间复杂度为 O(n)

	// 用于记录下标i左边的最大高度的dp
	leftDp := make([]int, len(height))
	leftDp[0] = height[0]
	for i := 1; i < len(height); i++ {
		leftDp[i] = max(height[i], leftDp[i-1])
	}
	// 用于记录下标i右边的最大高度的dp
	rightDp := make([]int, len(height))
	rightDp[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rightDp[i] = max(height[i], rightDp[i+1])
	}

	for i, h := range height {
		// 找出当前下标中，左、右高边的最小的那一个边(木桶原理，取短板)
		ret += min(leftDp[i], rightDp[i]) - h
	}
	return ret
}

// trapWithStack 基于单调递减栈实现
func trapWithStack(height []int) (ret int) {
	// 使用单调递减栈，具体思路是通过单调递减栈的性质，找出height图中的“凹陷处”
	var stack []int

	for i, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			topI := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 因为要从stack中获取topI的前一个下标，所以要判断stack的长度，以免下标溢出
			if len(stack) == 0 {
				break
			}
			// topI的左1位
			leftI := stack[len(stack)-1]
			// 计算宽度
			currWidth := i - leftI - 1
			// 计算高度差
			currHeight := min(height[leftI], h) - height[topI]

			ret += currWidth * currHeight
		}
		stack = append(stack, i)
	}
	return
}

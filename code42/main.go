package main

// trapWithDP 基于动态规划
// 基于x轴计算每个height的高度差，累加
func trapWithDP(height []int) (ret int) {
	// 用DP的话，时间复杂度为 O(n)；空间复杂度为 O(n)

	ldp := make([]int, len(height))
	ldp[0] = height[0]
	for i := 1; i < len(height); i++ {
		ldp[i] = min(ldp[i-1], height[i])
	}

	rdp := make([]int, len(height))
	e := len(height) - 1
	rdp[e] = height[e]
	for i := e - 1; i >= 0; i-- {
		rdp[i] = max(rdp[i+1], height[i])
	}

	for i, h := range height {
		ret += min(ldp[i], rdp[i]) - h
	}
	return
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

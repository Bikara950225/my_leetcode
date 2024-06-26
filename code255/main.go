package main

import "math"

// 这里使用了单调递减栈去探测左、右子树的边界；
// 当遍历每一个item时，最后一个从栈中pop的值为当前item的左子树启点（最后pop的数必定比当前值小）；
// 遍历数组时index越大说明越靠近右子树，如果发现右子树的值小于当前弹出的最小值（说明是相对于右子树的左子树），则说明这不是一棵正常的搜索树，因为搜索树必须右大于左
func verifyPreorder(preorder []int) bool {
	minN := math.MinInt32
	var stack []int
	for _, n := range preorder {
		if n < minN {
			return false
		}

		for len(stack) > 0 && n > stack[len(stack)-1] {
			minN = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, n)
	}
	return true
}

func main() {

}

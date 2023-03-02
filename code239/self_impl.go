package main

// 自己实现的，性能比较查但是可以通过
// 自己写的，性能比较慢
func maxSlidingWindow2(nums []int, k int) (ret []int) {
	upStack := []int{} // 单调递增栈
	m := map[int]int{} // 记录nums的值与nums下标的关系

	for i := 0; i < k; i++ {
		upStack = pushStack(upStack, m, nums[i], i)
	}
	ret = append(ret, upStack[0])

	for i := 1; ; i++ {
		j := i + k - 1
		if j >= len(nums) {
			break
		}
		// TODO 其实不用每次移动窗口都删除过期元素的
		// 将落后于窗口的元素删除
		pushItem := nums[i-1]
		pushIdx, ok := m[pushItem]
		if ok && i-1 == pushIdx { // 因为数组中的元素可能会有重复, 需要判断当前元素能不能出栈
			upStack = deleteItem(upStack, m, pushItem)
		}
		upStack = pushStack(upStack, m, nums[j], j)
		ret = append(ret, upStack[0])
	}
	return
}

func deleteItem(upStack []int, m map[int]int, pushItem int) []int {
	idx := len(upStack) - 1
	for ; idx >= 0; idx-- {
		if upStack[idx] == pushItem {
			break
		}
	}
	if idx < 0 {
		return upStack
	}

	upStack = append(upStack[:idx], upStack[idx+1:]...)
	delete(m, pushItem)
	return upStack
}

func pushStack(stack []int, m map[int]int, item, itemIdx int) []int {
	for len(stack) > 0 && item >= stack[len(stack)-1] {
		stack = stack[:len(stack)-1] // 出栈
	}
	stack = append(stack, item)
	m[item] = itemIdx
	return stack
}

// 这种是暴力解, 会超时
func maxSlidingWindow1(nums []int, k int) (ret []int) {
	for i := 0; i < len(nums); i++ {
		if i+k > len(nums) {
			break
		}
		maxItem := max(nums[i : i+k]...)
		ret = append(ret, maxItem)
	}
	return
}

func max(srcN ...int) int {
	ret := srcN[0]
	for i := 1; i < len(srcN); i++ {
		if srcN[i] > ret {
			ret = srcN[i]
		}
	}
	return ret
}

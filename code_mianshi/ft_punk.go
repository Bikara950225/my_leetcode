package code_mianshi

// 题目：
// 给定一个整形数组，找出所有满足如下要求的数: 所有它左边的数都比它小，所有它右边的数都比它大
// 要求： 请实现一个函数findNumber(giveArray)传指定数组返回结果数组，时间复杂度要求小于等于O(Xn), X为常数
// 例如： 传递数组 [1, 2, 4, 3, 10, 13, 15, 12, 16, 18, 17] , 数组中 1, 2, 10, 16 满足条件, 因此返回 [1,2,10,16]
func findNumber(src []int) (ret []int) {
	if len(src) <= 0 {
		return
	}

	// lDp记录当前下标到最左里面，最大的值
	lDp := make([]int, len(src))
	for i := 1; i < len(src); i++ {
		item := src[i]
		if src[lDp[i-1]] >= item {
			lDp[i] = lDp[i-1]
		} else {
			lDp[i] = i
		}
	}

	// rDp记录当前下标到最右里面，最小的值
	rDp := make([]int, len(src))
	end := len(src) - 1
	rDp[end] = end
	for i := end - 1; i >= 0; i-- {
		item := src[i]
		if src[rDp[i+1]] <= item {
			rDp[i] = rDp[i+1]
		} else {
			rDp[i] = i
		}
	}

	for i, item := range src {
		if lDp[i] == i && rDp[i] == i {
			ret = append(ret, item)
		}
	}
	return
}

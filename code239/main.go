package main

import (
	"log"
	"reflect"
)

func maxSlidingWindow(nums []int, k int) (ret []int) {
	// 这是一个单调递增的stack, 元素是nums的下标, 根据nums的值单调递增
	// 单调stack本来就可以保证其记录的元素单调递增, 所以不需要直接记录递增值，间接的记录它们的下标，这样做可以不用另开一个map，记录值与下标的关系!
	var q []int
	// 使用闭包, 把nums和q捕获到闭包中, 免去很多参数的传递, Slice的函数传递很麻烦
	pushStackFunc := func(idx int) {
		// nums[idx] == nums[q[len(q)-1]] 也要处理出栈, 更新记录的nums下标!!
		for len(q) > 0 && nums[idx] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, idx)
	}

	for i := 0; i < k; i++ {
		pushStackFunc(i)
	}
	ret = append(ret, nums[q[0]])

	for i := 1; ; i++ {
		j := i + k - 1
		if j >= len(nums) {
			break
		}
		pushStackFunc(j)
		if q[0] < i { // 指校验栈顶, 栈顶记录的下标不在窗口范围内的话就把队头删除
			q = q[1:]
		}
		ret = append(ret, nums[q[0]])
	}
	return
}

func main() {
	ret1 := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	expectRet1 := []int{3, 3, 5, 5, 6, 7}
	if !reflect.DeepEqual(ret1, expectRet1) {
		log.Fatalf("maxSlidingWindow error: ret1 != expectRet1: %+v != %+v", ret1, expectRet1)
	}

	ret2 := maxSlidingWindow([]int{1, -1}, 1)
	expectRet2 := []int{1, -1}
	if !reflect.DeepEqual(ret2, expectRet2) {
		log.Fatalf("maxSlidingWindow error: ret2 != expectRet2: %+v != %+v", ret2, expectRet2)
	}

	ret3 := maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3)
	expectRet3 := []int{3, 3, 2, 5}
	if !reflect.DeepEqual(ret3, expectRet3) {
		log.Fatalf("maxSlidingWindow error: ret3 != expectRet3: %+v != %+v", ret3, expectRet3)
	}
}

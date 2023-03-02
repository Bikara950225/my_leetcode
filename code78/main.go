package main

import (
	"fmt"
	"reflect"
)

func subsets(nums []int) [][]int {
	// subSet的大小；可以把subSet的排序类比为字节位的分布，比如：
	// nums: [1, 2, 3]
	// 0 0 0 => []
	// 0 1 0 => [2]
	// 1 1 0 => [2, 3]
	// 0 1 1 => [1, 2]
	// 1 1 1 => [1, 2, 3]
	subSetCount := 1 << len(nums)
	ret := make([][]int, 0, subSetCount)
	for i := 0; i < subSetCount; i++ {
		subSet := []int{}
		for j := 0; j < len(nums); j++ {
			// 通过位运算，判断当前下标对应值是否加入到subSet中
			// 比如当前i为5，即：0x101, 说明0、2下标的元素需要加入到subSet中
			// jBit是根据当前下标计算出的状态位结果, 用来和i进行与运算，即可判断出当前下标是否使用
			jBit := 1 << j
			if jBit&i == jBit {
				subSet = append(subSet, nums[j])
			}
		}
		ret = append(ret, subSet)
	}
	return ret
}

func main() {
	ret1 := subsets([]int{1, 2, 3})
	expect1 := [][]int{
		{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3},
	}
	if !reflect.DeepEqual(ret1, expect1) {
		panic(fmt.Errorf("code78 error, not expect: %+v", ret1))
	}
}

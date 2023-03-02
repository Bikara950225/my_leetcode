package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 0 {
		return nil
	}
	sortIntervals(intervals, 0, len(intervals)-1)

	var ret [][]int
	for i := 0; i < len(intervals); i++ {
		// 处理intervals第一个元素时
		if i == 0 {
			ret = append(ret, intervals[0])
			continue
		}
		currInterval := intervals[i]
		backInterval := ret[len(ret)-1]

		// 当前数组的范围不在融合数组的范围内，不进行融合，作为新的数组集入队到ret结果集中
		if backInterval[1] < currInterval[0] {
			ret = append(ret, currInterval)
			continue
		}

		// merge
		// 当前数组开辟了新的范围，更新对比数组的范围
		if backInterval[1] < currInterval[1] {
			backInterval[1] = currInterval[1]
		}
	}
	return ret
}

func sortIntervals(src [][]int, begin, end int) {
	if begin >= end {
		return
	}

	randIndex := (rand.Int() % (end - begin + 1)) + begin
	mark := src[randIndex][0]
	swapIntervals(src, end, randIndex)

	i := begin - 1
	for j := begin; j < end; j++ {
		if src[j][0] < mark {
			i++
			swapIntervals(src, i, j)
		}
	}
	i += 1
	swapIntervals(src, i, end)

	sortIntervals(src, begin, i-1)
	sortIntervals(src, i+1, end)
}

func swapIntervals(src [][]int, i1, i2 int) {
	if i1 == i2 {
		return
	}
	src[i1], src[i2] = src[i2], src[i1]
}

func main() {
	srcIntervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	ret := merge(srcIntervals)
	expectRet := [][]int{
		{1, 6},
		{8, 10},
		{15, 18},
	}
	if !reflect.DeepEqual(ret, expectRet) {
		panic(fmt.Errorf("code56 error, not expect result: %+v", ret))
	}
}

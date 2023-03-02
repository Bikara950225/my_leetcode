package main

import (
	"fmt"
	"my_leetcode/internal/sort"
	"reflect"
)

func threeSum(nums []int) (ret [][]int) {
	if len(nums) <= 0 {
		return
	}

	sort.QuickSort(nums)

	for first := 0; first < len(nums); first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		trd := len(nums) - 1
		for second := first + 1; second < len(nums); second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			subSum := nums[first] + nums[second]
			for second < trd && subSum+nums[trd] > 0 {
				trd--
			}

			if second >= trd {
				break
			}
			if subSum+nums[trd] == 0 {
				ret = append(ret, []int{
					nums[first], nums[second], nums[trd],
				})
			}
		}
	}
	return
}

func main() {
	ret := threeSum([]int{-1, 0, 1, 2, -1, -4})
	expectRet := [][]int{
		{-1, -1, 2},
		{-1, 0, 1},
	}
	if !reflect.DeepEqual(ret, expectRet) {
		panic(fmt.Errorf("not match expectRet: %+v", ret))
	}
}

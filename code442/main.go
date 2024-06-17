package main

import (
	"fmt"
	"reflect"
	"sort"
)

//func findDuplicates(nums []int) (ret []int) {
//	for i := range nums {
//		for nums[i] != nums[nums[i]-1] {
//			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
//		}
//	}
//	for i, num := range nums {
//		if num-1 != i {
//			ret = append(ret, num)
//		}
//	}
//	return
//}

func findDuplicates(nums []int) (ret []int) {
	for i := range nums {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i, num := range nums {
		if i+1 != num {
			ret = append(ret, num)
		}
	}
	return
}

func main() {
	ret := findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1})
	sort.SliceStable(ret, func(i, j int) bool {
		return ret[i] < ret[j]
	})
	if !reflect.DeepEqual(ret, []int{2, 3}) {
		panic(fmt.Errorf("want != got: %v != %v", []int{2, 3}, ret))
	}
}

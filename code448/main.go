package main

import (
	"fmt"
	"reflect"
)

func findDisappearedNumbers(nums []int) (ret []int) {
	for i := range nums {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i, num := range nums {
		if i != num-1 {
			ret = append(ret, i+1)
		}
	}
	return
}

func main() {
	src1 := []int{4, 3, 2, 7, 8, 2, 3, 1}
	ret1 := findDisappearedNumbers(src1)
	want1 := []int{5, 6}
	if !reflect.DeepEqual(ret1, want1) {
		panic(fmt.Errorf("code448 error, got: %v but: %v", ret1, want1))
	}
}

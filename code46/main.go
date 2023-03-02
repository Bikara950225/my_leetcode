package main

import (
	"fmt"
	"my_leetcode/internal/empty"
	"reflect"
)

var (
	ret [][]int
	sub []int
	mark = map[int]empty.Empty{}
)

func permute(nums []int) [][]int {
	handle(nums)

	return ret
}

func handle(nums []int) {
	if len(sub) == len(nums) {
		copySub := make([]int, len(sub))
		copy(copySub, sub)
		ret = append(ret, copySub)
	}

	for _, item := range nums {
		if _, ok := mark[item]; ok {
			continue
		}

		sub = append(sub, item)
		mark[item] = empty.Empty{}

		handle(nums)

		sub = sub[:len(sub)-1]
		delete(mark, item)
	}
}

func main() {
	defer func(){
		ret = nil
		sub = nil
		mark = map[int]empty.Empty{}
	}()

	getRet := permute([]int{1, 2, 3})
	expectRet := [][]int{
		{1,2,3},
		{1,3,2},
		{2,1,3},
		{2,3,1},
		{3,1,2},
		{3,2,1},
	}
	if !reflect.DeepEqual(getRet, expectRet) {
		panic(fmt.Errorf("code46 error: not expect result: \n%+v", getRet))
	}
}

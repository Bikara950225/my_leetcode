package main

import (
	"fmt"
	"reflect"
)

func permute(nums []int) (ret [][]int) {
	mark := map[int]struct{}{}
	subRet := make([]int, 0, len(nums))

	var ff func()
	ff = func() {
		if len(subRet) == len(nums) {
			cpRet := make([]int, len(subRet))
			copy(cpRet, subRet)
			ret = append(ret, cpRet)
			return
		}

		for i, num := range nums {
			if _, ok := mark[i]; ok {
				continue
			}

			subRet = append(subRet, num)
			mark[i] = struct{}{}
			ff()
			subRet = subRet[:len(subRet)-1]
			delete(mark, i)
		}
	}
	ff()
	return
}

func main() {
	getRet := permute([]int{1, 2, 3})
	expectRet := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}
	if !reflect.DeepEqual(getRet, expectRet) {
		panic(fmt.Errorf("code46 error: not expect result: \n%+v", getRet))
	}
}

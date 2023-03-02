package main

import (
	"fmt"
	"reflect"
)

func spiralOrder(matrix [][]int) []int {
	var ret []int

	top := 0
	down := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1

	for top <= down && left <= right {
		// 顶边遍历
		for l := left; l <= right; l++ {
			ret = append(ret, matrix[top][l])
		}
		top += 1
		if top > down || left > right {
			break
		}

		// 右边遍历
		for t := top; t <= down; t++ {
			ret = append(ret, matrix[t][right])
		}
		right -= 1
		if top > down || left > right {
			break
		}

		// 底边遍历
		for r := right; r >= left; r-- {
			ret = append(ret, matrix[down][r])
		}
		down -= 1
		if top > down || left > right {
			break
		}

		// 左边遍历
		for d := down; d >= top; d-- {
			ret = append(ret, matrix[d][left])
		}
		left += 1
	}

	return ret
}

func main() {
	ret := spiralOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
	expectRet := []int{
		1, 2, 3, 6, 9, 8, 7, 4, 5,
	}

	if !reflect.DeepEqual(ret, expectRet) {
		panic(fmt.Errorf("code54 error: not expect result: \n%+v", ret))
	}
}

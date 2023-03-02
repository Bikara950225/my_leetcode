package main

import (
	"fmt"
	"my_leetcode/internal/tools"
)

var (
	ret = 0
)

func maximalSquare(matrix [][]byte) int {
	defer func() {
		ret = 0
	}()

	// matrix每个列与行都保证相等
	yLen := len(matrix)
	xLen := len(matrix[0])
	dp := tools.NewInt2DP(xLen, yLen)

	for y := 0; y < yLen; y++ {
		for x := 0; x < xLen; x++ {
			if matrix[y][x] == '0' {
				continue
			}

			if x == 0 || y == 0 {
				// 边缘的节点，不用计算它之前的dp
				// matrix不是'0'就是'1'，和 '0' 取差值可以直接得到答案
				dp[y][x] = int(matrix[y][x] - '0')
			} else {
				// 取当前节点的左、上、左上节点的dp，取最小值 + 1
				dp[y][x] = tools.IntMin(dp[y-1][x], dp[y][x-1], dp[y-1][x-1]) + 1
			}
			area := dp[y][x] * dp[y][x]
			if area > ret {
				ret = area
			}
		}
	}
	return ret
}

func main() {
	// [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
	matrix1 := tools.NewByteMatrix(
		[]byte{'1','0','1','0','0'},
		[]byte{'1','0','1','1','1'},
		[]byte{'1','1','1','1','1'},
		[]byte{'1','0','0','1','0'},
	)
	ret1 := maximalSquare(matrix1)
	expectRet1 := 4
	if ret1 != expectRet1 {
		panic(fmt.Errorf("code221 error, not expect result: %d\n", ret1))
	}
}

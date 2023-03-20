package main

import "fmt"

func maximalRectangle(matrix [][]byte) (ret int) {
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '0' {
				continue
			}
			if j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i][j-1] + 1
			}
			if dp[i][j] > ret {
				ret = dp[i][j]
			}

			tmp := dp[i][j]
			for z := i - 1; z >= 0; z-- {
				if matrix[z][j] == '0' {
					break
				}
				tmp = min(tmp, dp[z][j])
				area := tmp * (i - z + 1)
				if area > ret {
					ret = area
				}
			}
		}
	}
	return
}

func min(src1, src2 int) int {
	if src1 < src2 {
		return src1
	}
	return src2
}

func main() {
	matrix1 := [][]byte{
		[]byte("10100"),
		[]byte("10111"),
		[]byte("11111"),
		[]byte("10010"),
	}
	ret1 := maximalRectangle(matrix1)
	if ret1 != 6 {
		panic(fmt.Errorf("code85 error: ret != 6: %d", ret1))
	}
}

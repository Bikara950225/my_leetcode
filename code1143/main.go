package main

import "fmt"

func longestCommonSubsequence(text1, text2 string) int {
	// 二维DP，dp[x][y]意思是 text1[0:x] 和 text[0:y] 的重复子序列的长度
	dp := make([][]int, len(text1)+1)
	for i := 0; i <= len(text1); i++ {
		dp[i] = make([]int, len(text2)+1)
	}

	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

func max(src1, src2 int) int {
	if src1 > src2 {
		return src1
	}
	return src2
}

func main() {
	ret1 := longestCommonSubsequence("abcde", "ace")
	if ret1 != 3 {
		panic(fmt.Errorf("code1143 error: ret1 != 3: %d", ret1))
	}
}

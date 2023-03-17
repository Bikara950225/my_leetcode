package main

func max(src1, src2 int) int {
	if src1 > src2 {
		return src1
	}
	return src2
}

// 最长回文子串
// [1,1,1,2,1] => [1,1,1,1]
func longestPalindromeSubseq(s string) (ret int) {
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
		for j := range dp[i] {
			dp[i][j] = 1
		}
	}
	ret = 1

	for i := 1; i < len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if s[i] == s[j] {
				if i == j+1 {
					dp[j][i] = 2
				} else {
					dp[j][i] = dp[j+1][i-1] + 2
				}
			} else {
				dp[j][i] = max(dp[j+1][i], dp[j][i-1])
			}
		}
	}
	return dp[0][len(dp)-1]
}

func main() {
	ret1 := longestPalindromeSubseq("11121")
	if ret1 != 4 {
		panic("stupid")
	}
}

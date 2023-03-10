package main

import "fmt"

// bbbab => bbbb
// bab => bab
// cbbd => bb
func longestPalindromeSubseq(s string) (ret int) {
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
		for j := range dp[i] {
			dp[i][j] = 1
		}
	}

	for i := 1; i < len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if s[i] == s[j] {
				if j+1 == i {
					dp[j][i] = 2
				} else {
					dp[j][i] = dp[j+1][i-1] + 2
				}
				continue
			}
			dp[j][i] = max(dp[j+1][i], dp[j][i-1])
		}
	}
	return dp[0][len(dp)-1]
}

func max(src1, src2 int) int {
	if src1 > src2 {
		return src1
	}
	return src2
}

func main() {
	ret1 := longestPalindromeSubseq("bbbab")
	if ret1 != 4 {
		panic(fmt.Errorf("code516 error, ret1 != 4: %d", ret1))
	}

	ret2 := longestPalindromeSubseq("cbbd")
	if ret2 != 2 {
		panic(fmt.Errorf("code516 error, ret2 != 2: %d", ret2))
	}

	ret3 := longestPalindromeSubseq("bab")
	if ret3 != 3 {
		panic(fmt.Errorf("code516 error, ret3 != 3: %d", ret3))
	}
}

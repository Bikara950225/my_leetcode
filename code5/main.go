package main

import "fmt"

func longestPalindrome(s string) string {
	// 该dp[x][y] 表示 从下标 x 到 y 的子字符串是否为回文数
	dp := make([][]bool, 0, len(s))
	for i := 0; i < len(s); i++ {
		dp = append(dp, make([]bool, len(s)))
	}

	var ret string
	for step := 0; step < len(s); step++ {
		for begin := 0; begin < len(s); begin++ {
			end := begin + step
			if end >= len(s) {
				break
			}

			if begin == end {
				dp[begin][end] = true
			} else if begin+1 == end {
				dp[begin][end] = s[begin] == s[end]
			} else {
				dp[begin][end] = dp[begin+1][end-1] && (s[begin] == s[end])
			}

			if dp[begin][end] && end-begin+1 > len(ret) {
				ret = s[begin : end+1]
			}
		}
	}
	return ret
}

func main() {
	src := "babad"

	ret := longestPalindrome(src)
	expectRet := "bab"
	if ret != expectRet {
		panic(fmt.Errorf("code5 error, not expect result: %s", ret))
	}
}

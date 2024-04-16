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

func longestPalindrome2(s string) (ret string) {
	dp := make([][]bool, 0, len(s))
	for range s {
		dp = append(dp, make([]bool, len(s)))
	}

	for step := 0; step < len(s); step++ {
		for i := range s {
			j := i + step
			if j >= len(s) {
				break
			}
			if i == j {
				dp[i][j] = true
			} else if i+1 == j {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = (s[i] == s[j]) && dp[i+1][j-1]
			}
			if dp[i][j] {
				if j-i+1 > len(ret) {
					ret = s[i : j+1]
				}
			}
		}
	}
	return
}

func method(src string) (ret string) {
	dp := make([][]bool, len(src))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(src))
	}

	for step := 0; step < len(src); step++ {
		for i := 0; ; i++ {
			j := i + step
			if j >= len(src) {
				break
			}

			if i == j {
				dp[i][j] = true
			} else if i+1 == j {
				dp[i][j] = src[i] == src[j]
			} else {
				dp[i][j] = dp[i+1][j-1] && src[i] == src[j]
			}

			if dp[i][j] {
				if j-i+1 > len(ret) {
					ret = src[i : j+1]
				}
			}
		}
	}
	return
}

func main() {
	src := "babad"

	ret := longestPalindrome2(src)
	expectRet := "bab"
	if ret != expectRet {
		panic(fmt.Errorf("code5 error, not expect result: %s", ret))
	}

	src = "cbbd"
	ret = longestPalindrome2(src)
	expectRet = "bb"
	if ret != expectRet {
		panic(fmt.Errorf("code5 error, not expect result: %s", ret))
	}
}

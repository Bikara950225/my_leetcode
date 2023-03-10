package main

import "fmt"

func longestValidParentheses(s string) (ret int) {
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	for step := 1; step < len(s); step++ {
		validate := false
		for i := 0; ; i++ {
			j := i + step
			if j >= len(s) {
				break
			}

			if i+1 == j {
				dp[i][j] = s[i] == '(' && s[j] == ')'
				goto checkValidate
			}

			if s[i] == '(' && s[j] == ')' {
				dp[i][j] = check(i, j, dp)
			}

		checkValidate:
			if dp[i][j] {
				validate = true
			}
		}

		if validate {
			if step > ret {
				ret = step + 1
			}
		}
	}
	return
}

func check(i, j int, dp [][]bool) bool {
	if dp[i+1][j-1] {
		return true
	}

	ib, ie := i, i+1
	for ; ie < j; ie++ {
		if dp[ib][ie] {
			break
		}
	}
	if ie == j {
		return false
	}

	jb, je := j-1, j
	for ; jb > i; jb-- {
		if dp[jb][je] {
			break
		}
	}
	if jb == i {
		return false
	}
	if ie+1 > jb-1 {
		return true
	}
	return dp[ie+1][jb-1]
}

func main() {
	ret1 := longestValidParentheses("(())")
	if ret1 != 4 {
		panic(fmt.Sprintf("ret1 result not 4: %d", ret1))
	}
}

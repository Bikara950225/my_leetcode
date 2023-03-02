package main

import "fmt"

func lengthOfLIS(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	ret := 1
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		listCount := 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] && listCount < dp[j]+1 {
				listCount = dp[j] + 1
			}
		}

		dp[i] = listCount
		if dp[i] > ret {
			ret = dp[i]
		}
	}
	return ret
}

func main() {
	ret := lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})
	expectRet := 4 // {2,3,7,101}
	if ret != expectRet {
		panic(fmt.Errorf("code300 error, not expect result: %d != %d", ret, expectRet))
	}

	ret = lengthOfLIS([]int{7, 7, 7, 7, 7})
	expectRet = 1 // {7}
	if ret != expectRet {
		panic(fmt.Errorf("code300 error, not expect result: %d != %d", ret, expectRet))
	}

	ret = lengthOfLIS([]int{0})
	expectRet = 1 // {0}
	if ret != expectRet {
		panic(fmt.Errorf("code300 error, not expect result: %d != %d", ret, expectRet))
	}
}

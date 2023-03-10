package main

import "fmt"

type mark struct {
	length int
	count  int
}

// [1,3,5,4,7] => [1,3,4,7]、[1,3,5,7] = 2
// [1,3,5,4,7,6] => [1,3,4,6]、[1,3,4,7]、[1,3,5,7]、[1,3,5,6]
// [5,5,5,5,5] => 5 * [5] = 5
// 1, 2, 4, 3, 5, 4, 7, 2
func findNumberOfLIS(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	dp := make([]mark, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = mark{1, 1}
	}
	ret, currLen := 1, 1
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				if dp[j].length+1 > dp[i].length {
					dp[i].length = dp[j].length + 1
					dp[i].count = dp[j].count
				} else if dp[j].length+1 == dp[i].length {
					dp[i].count += dp[j].count
				}
			}
		}
		if dp[i].length > currLen {
			currLen = dp[i].length
			ret = dp[i].count
		} else if dp[i].length == currLen {
			ret += dp[i].count
		}
	}
	return ret
}

func main() {
	ret1 := findNumberOfLIS([]int{1, 3, 5, 4, 7, 6})
	if ret1 != 4 {
		panic(fmt.Errorf("code673 error, ret1 != 4: %d", ret1))
	}

	ret2 := findNumberOfLIS([]int{1, 3, 5, 4, 7})
	if ret2 != 2 {
		panic(fmt.Errorf("code673 error, ret2 != 2: %d", ret2))
	}

	ret3 := findNumberOfLIS([]int{5, 5, 5, 5, 5})
	if ret3 != 5 {
		panic(fmt.Errorf("code673 error, ret3 != 5: %d", ret3))
	}

	ret4 := findNumberOfLIS([]int{1, 2, 4, 3, 5, 4, 7, 2})
	if ret4 != 3 {
		panic(fmt.Errorf("code673 error, ret4 != 3: %d", ret4))
	}

	ret5 := findNumberOfLIS([]int{3, 1, 2})
	if ret5 != 1 {
		panic(fmt.Errorf("code673 error, ret5 != 1： %d", ret5))
	}
}

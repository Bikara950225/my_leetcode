package main

import "fmt"

func maxSubArray(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	dp := nums[0]
	ret := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i]+dp {
			dp = nums[i]
		} else {
			dp += nums[i]
		}
		if dp > ret {
			ret = dp
		}
	}
	return ret
}

func main() {
	maxSum := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	expectSum := 6 // 4,-1,2,1

	if maxSum != expectSum {
		panic(fmt.Errorf("not expectSum: %d", maxSum))
	}
}

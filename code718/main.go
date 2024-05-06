package main

import "fmt"

//func findLength(nums1 []int, nums2 []int) (ret int) {
//	dp := make([][]int, len(nums1))
//	for i := range dp {
//		dp[i] = make([]int, len(nums2))
//	}
//
//	for i := 0; i < len(nums1); i++ {
//		for j := 0; j < len(nums2); j++ {
//			if nums1[i] == nums2[j] {
//				if i == 0 || j == 0 {
//					dp[i][j] = 1
//				} else {
//					dp[i][j] = dp[i-1][j-1] + 1
//				}
//			} else {
//				dp[i][j] = 0
//			}
//			if dp[i][j] > ret {
//				ret = dp[i][j]
//			}
//		}
//	}
//	return
//}

func findLength(nums1 []int, nums2 []int) (ret int) {
	dp := make([][]int, len(nums1))
	for i := range dp {
		dp[i] = make([]int, len(nums2))
	}

	for i := range nums1 {
		for j := range nums2 {

			if nums1[i] == nums2[j] {
				pre := 0
				if i-1 >= 0 && j-1 >= 0 {
					pre = dp[i-1][j-1]
				}
				dp[i][j] = pre + 1
			}
			if dp[i][j] > ret {
				ret = dp[i][j]
			}
		}
	}

	return
}

func main() {
	ret1 := findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7})
	if ret1 != 3 {
		panic(fmt.Sprintf("code718 error, ret1 != 3: %d", ret1))
	}
}

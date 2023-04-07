package main

import "fmt"

func maxSubArrayLen(nums []int, k int) (ret int) {
	var preSum int
	preSumMap := map[int]int{0: -1}

	for i, v := range nums {
		preSum += v
		ii, ok := preSumMap[preSum-k]
		if ok {
			if i-ii > ret {
				ret = i - ii
			}
		}
		// 只添加第一次出现的，从左到右遍历，先记录的前缀和说明index更小，那么其子数组的长度更大
		if _, ok = preSumMap[preSum]; !ok {
			preSumMap[preSum] = i
		}
	}
	return
}

// 这个因为是接近于O(N*logN)，所以会超时。。。
func maxSubArrayLen2(nums []int, k int) (ret int) {
	preSum := make([]int, len(nums))
	for i, v := range nums {
		if i == 0 {
			preSum[i] = nums[i]
		} else {
			preSum[i] = preSum[i-1] + v
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if i+1 < ret {
			break
		}
		if preSum[i] == k && i+1 > ret {
			ret = i + 1
			continue
		}

		for j := 0; j < i; j++ {
			subSum := preSum[i] - preSum[j]
			if subSum == k && i-j > ret {
				ret = i - j
				break
			}
		}
	}
	return
}

func main() {
	// 1, -1, 5, -2
	ret1 := maxSubArrayLen([]int{1, -1, 5, -2, 3}, 3)
	if ret1 != 4 {
		panic(fmt.Errorf("code325 error, ret1 != 4: %d", ret1))
	}

	// -1, 2
	ret2 := maxSubArrayLen([]int{-2, -1, 2, 1}, 1)
	if ret2 != 2 {
		panic(fmt.Errorf("code325 error, ret2 != 2: %d", ret2))
	}
}

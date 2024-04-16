package main

import "fmt"

// 时间复杂度: O(N)
func minSubArrayLen(target int, nums []int) int {
	b, e := 0, 0
	sum, ret := 0, -1
	for e < len(nums) {
		sum += nums[e]
	skipSum:
		if sum < target {
			e++
			continue
		}

		// sum >= target
		l := e - b + 1
		if l < ret || ret == -1 {
			ret = l
		}
		sum -= nums[b]
		b += 1

		if b > e {
			e = b
		} else {
			goto skipSum
		}
	}
	if ret == -1 {
		return 0
	}
	return ret
}

func minSubArrayLen2(target int, src []int) (ret int) {
	i, j := 0, 0
	sum := 0
	for i < len(src) && j < len(src) {
		sum += src[j]
	gto:
		if sum >= target {
			l := j - i + 1
			if l < ret || ret == 0 {
				ret = l
			}

			sum -= src[i]
			i++
			if i > j {
				j = i
			} else {
				goto gto
			}

		} else {
			j++
		}
	}
	return
}

// 只用了前缀和, 勉强可以过
// 时间复杂度: O(N*logN)
func minSubArrayLen1(target int, nums []int) int {
	preSum := make([]int, len(nums))
	for i, v := range nums {
		if i == 0 {
			preSum[i] = v
		} else {
			preSum[i] = preSum[i-1] + v
		}
	}
	if nums[0] >= target {
		return 1
	}

	ret := -1
	for i := len(nums) - 1; i > 0; i-- {
		if preSum[i] < target {
			// 当前 0到i 的总和也小于target，其子和、后续的总和肯定小于target，所以直接break
			break
		}
		for j := i - 1; j >= 0; j-- {
			subSum := preSum[i] - preSum[j]
			l := i - j
			if subSum >= target {
				if l < ret || ret == -1 {
					ret = l
				}
				break
			}
		}
		if preSum[i] >= target && (i+1 < ret || ret == -1) {
			ret = i + 1
		}
	}

	if ret == -1 {
		return 0
	}
	return ret
}

func main() {
	ret1 := minSubArrayLen2(7, []int{2, 3, 1, 2, 4, 3})
	if ret1 != 2 {
		panic(fmt.Errorf("code209 error: ret1 != 2: %d", ret1))
	}

	ret2 := minSubArrayLen2(4, []int{1, 4, 4, 4})
	if ret2 != 1 {
		panic(fmt.Errorf("code209 error: ret2 != 1: %d", ret2))
	}
}

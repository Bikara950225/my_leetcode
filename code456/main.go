package main

// 3,5,0,3,4
// 1,2,3,4
// 3,1,4,2
// 1,3,2,4
func find132pattern(nums []int) bool {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i, num := range nums[1:] {
		dp[i+1] = min(num, dp[i])
	}

	mp := map[int]int{}
	var stack []int
	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]
		lastMin := i
		for len(stack) > 0 && num > nums[stack[len(stack)-1]] {
			lastMin = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		mp[i] = lastMin
		stack = append(stack, i)
	}

	for i, num := range nums {
		left := dp[i]
		right := nums[mp[i]]
		if num > left && num > right && left < right {
			return true
		}
	}
	return false
}

func find132patternLoop(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				continue
			}
			for z := j + 1; z < len(nums); z++ {
				if nums[j] > nums[z] && nums[z] > nums[i] {
					return true
				}
			}
		}
	}
	return false
}

func main() {

}

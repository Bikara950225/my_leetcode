package main

func firstMissingPositive(nums []int) int {
	for i := range nums {
		for {
			// 限定nums的元素范围在[1, len(nums)]
			// item小于1丢弃为0
			// item大于len(nums)丢弃为0
			if nums[i] < 1 || nums[i] > len(nums) {
				nums[i] = 0
				break
			}
			if nums[i] == nums[nums[i]-1] {
				break
			}
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i, num := range nums {
		if i+1 != num {
			return i + 1
		}
	}
	return len(nums) + 1
}

func main() {
}

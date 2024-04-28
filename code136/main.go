package code136

func singleNumber(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		sum ^= nums[i]
	}
	return sum
}

package main

// 思路过程：
// 1. 使用Set记录元素的次数，超出 len(nums)/2时则返回该元素的值，时间复杂度为O(n)，空间复杂度为O(n)
// 2. 要求空间复杂度为O(1)的话，则可以先考虑排序，然后取len(nums)/2的下标，时间复杂度变为O(N*logN)
// 3. 再要求缩短时间复杂度时，才采用摩尔投票法
func majorityElement(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	count := 1
	ret := nums[0]
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			count = 1
			ret = nums[i]
			continue
		}

		if nums[i] == ret {
			count += 1
		} else {
			count -= 1
		}
	}
	return ret
}

func main() {

}

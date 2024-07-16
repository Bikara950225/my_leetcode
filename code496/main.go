package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	return nextGreaterElement2(nums1, nums2)
}

func nextGreaterElement1(nums1 []int, nums2 []int) []int {
	ret := make([]int, len(nums1))
	for i := range ret {
		ret[i] = -1
	}

	num2map := map[int]int{}
	for i, num := range nums2 {
		num2map[num] = i
	}

	for i, num := range nums1 {
		for j := num2map[num] + 1; j < len(nums2); j++ {
			if nums2[j] > num {
				ret[i] = nums2[j]
				break
			}
		}
	}
	return ret
}

func nextGreaterElement2(nums1, nums2 []int) []int {
	mp := map[int]int{}
	var stack []int
	for i := len(nums2) - 1; i >= 0; i-- {
		num := nums2[i]
		for len(stack) > 0 && num > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			mp[num] = stack[len(stack)-1]
		} else {
			mp[num] = -1
		}
		stack = append(stack, num)
	}

	ret := make([]int, 0, len(nums1))
	for _, num := range nums1 {
		ret = append(ret, mp[num])
	}
	return ret
}

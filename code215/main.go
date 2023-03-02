package main

import (
	"fmt"
	"math/rand"
)

func findKthLargest(nums []int, k int) int {
	for i := len(nums) >> 1; i >= 0; i-- {
		maxHeap(nums, i, len(nums))
	}

	for i := 0; i < k; {
		nums[0], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[0]
		i++
		maxHeap(nums, 0, len(nums)-i)
	}
	return nums[len(nums)-k]
}

func maxHeap(nums []int, b, e int) {
	for b<<1+1 < e {
		largest := b

		left := b<<1 + 1
		if left < e && nums[left] > nums[largest] {
			largest = left
		}
		right := b<<1 + 2
		if right < e && nums[right] > nums[largest] {
			largest = right
		}

		if largest != b {
			nums[b], nums[largest] = nums[largest], nums[b]
			b = largest
		} else {
			break
		}
	}
}

func main() {
	ret := findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
	if ret != 5 {
		panic(fmt.Errorf("error, expect result is 5, but %d", ret))
	}
	rand.Int()
}

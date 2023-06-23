package main

import "fmt"

// 思路过程：
// 1. 使用Set记录元素的次数，超出 len(nums)/2时则返回该元素的值，时间复杂度为O(n)，空间复杂度为O(n)
// 2. 要求空间复杂度为O(1)的话，则可以先考虑排序，然后取len(nums)/2的下标，时间复杂度变为O(N*logN)
// 3. 再要求缩短时间复杂度时，才采用摩尔投票法
func majorityElement1(nums []int) int {
	// 采用count计数，比栈记录更加省内存
	if len(nums) <= 0 {
		return -1
	}
	count := 1
	mark := nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num == mark {
			count++
			continue
		}
		if count <= 0 {
			count = 1
			mark = num
		} else {
			count--
		}
	}
	return mark
}

func majorityElement2(nums []int) int {
	// 自己想的用stack解决，比较耗空间
	stack := make([]int, 0, len(nums))
	for _, num := range nums {
		if len(stack) > 0 && num != stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, num)
		}
	}
	return stack[len(stack)-1]
}

func main() {
	ret1 := majorityElement1([]int{2, 2, 2, 1, 1})
	if ret1 != 2 {
		panic(fmt.Sprintf("code169 error, ret1 != 2: %d\n", ret1))
	}

	ret2 := majorityElement2([]int{2, 2, 2, 1, 1})
	if ret2 != 2 {
		panic(fmt.Sprintf("code169 error, ret2 != 2: %d\n", ret2))
	}
}

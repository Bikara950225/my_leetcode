package main

import "fmt"

func subarraySum(nums []int, k int) (ret int) {
	for i := 0; i < len(nums); i++ {
		var sum int
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				ret += 1
			}
		}
	}
	return
}

func main() {
	ret1 := subarraySum([]int{1, 1, 1}, 2)
	if ret1 != 2 {
		panic(fmt.Errorf("code560 error, ret1 != 2: %d", ret1))
	}

	fmt.Println("code560 Done")
}

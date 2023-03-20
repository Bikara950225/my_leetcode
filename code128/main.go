package main

import "fmt"

func longestConsecutive(nums []int) (ret int) {
	numSet := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		numSet[num] = struct{}{}
	}

	for _, num := range nums {
		if _, ok := numSet[num-1]; ok {
			continue
		}
		subRet := 0
		for {
			if _, ok := numSet[num]; !ok {
				break
			}
			subRet++
			num++
		}
		if subRet > ret {
			ret = subRet
		}
	}
	return
}

func main() {
	ret1 := longestConsecutive([]int{1, 4, 3, 5, 6, 2}) // 1，2，3，4，5，6
	if ret1 != 6 {
		panic(fmt.Errorf("code128 error, not expect ret1: %d", ret1))
	}
}

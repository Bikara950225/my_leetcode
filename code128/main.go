package main

import "fmt"

func longestConsecutive(nums []int) (ret int) {
	if len(nums) <= 0 {
		return
	}
	ret = 1

	m := map[int]struct{}{}
	for _, num := range nums {
		m[num] = struct{}{}
	}

	for _, num := range nums {
		if _, ok := m[num-1]; ok {
			continue
		}

		loopRet := 1
		for currNum := num+1; ; currNum++ {
			if _, ok := m[currNum]; !ok {
				if loopRet > ret {
					ret = loopRet
				}
				break
			}
			loopRet += 1
		}
	}
	return
}

func main() {
	ret1 := longestConsecutive([]int{1, 4, 3, 5, 6, 2})
	if ret1 != 6 {
		panic(fmt.Errorf("code128 error, not expect ret1: %d", ret1))
	}
}

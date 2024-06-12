package main

import "fmt"

func nextGreaterElement(n int) int {
	src := []int{}
	nn := n
	for nn > 0 {
		src = append(src, nn%10)
		nn /= 10
	}
	reserve(src)

	downIdx := -1
	for i := len(src) - 1; i > 0; i-- {
		if src[i] > src[i-1] {
			downIdx = i - 1
			break
		}
	}
	if downIdx < 0 {
		return -1
	}

	for i := len(src) - 1; i > downIdx; i-- {
		if src[i] > src[downIdx] {
			src[i], src[downIdx] = src[downIdx], src[i]
			break
		}
	}
	reserve(src[downIdx+1:])

	return recover(src, int32(n))
}

func recover(src []int, n int32) int {
	var ret int32
	for i, item := range src {
		if i == 0 {
			ret = int32(item)
		} else {
			ret = 10*ret + int32(item)
		}
	}

	if ret < n {
		return -1
	}
	return int(ret)
}

func reserve(src []int) {
	i, j := 0, len(src)-1
	for i < j {
		src[i], src[j] = src[j], src[i]
		i++
		j--
	}
}

func main() {
	ret1 := nextGreaterElement(101)
	if ret1 != 110 {
		panic(fmt.Errorf("case1 error, ret1 != 110: %d", ret1))
	}

	ret2 := nextGreaterElement(21)
	if ret2 != -1 {
		panic(fmt.Errorf("case2 error, ret2 != -1: %d", ret2))
	}
}

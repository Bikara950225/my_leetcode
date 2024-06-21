package main

import "fmt"

func nextGreaterElement(n int) int {
	// n to slice
	nn := n
	var nl []uint8
	for nn > 0 {
		nl = append(nl, uint8(nn%10))
		nn /= 10
	}
	reverse(nl)

	downI := -1
	for i := len(nl) - 1; i > 0; i-- {
		if nl[i] > nl[i-1] {
			downI = i - 1
			break
		}
	}
	if downI == -1 {
		return -1
	}

	for i := len(nl) - 1; i > downI; i-- {
		if nl[i] > nl[downI] {
			nl[i], nl[downI] = nl[downI], nl[i]
			break
		}
	}
	reverse(nl[downI+1:])

	return recoverNum(nl)
}

func reverse(src []uint8) {
	i, j := 0, len(src)-1
	for i < j {
		src[i], src[j] = src[j], src[i]
		i++
		j--
	}
}

func recoverNum(nl []uint8) (ret int) {
	if len(nl) <= 0 {
		return
	}

	ret = int(nl[0])
	for _, item := range nl[1:] {
		ret = 10*ret + int(item)
	}
	return
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

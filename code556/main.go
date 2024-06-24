package main

import (
	"fmt"
	"math"
)

func nextGreaterElement(n int) int {
	// n to slice
	nn := n
	var nl []uint8
	for nn > 0 {
		nl = append(nl, uint8(nn%10))
		nn /= 10
	}
	// nl为小端

	di := -1
	for i := 0; i < len(nl)-1; i++ {
		if nl[i] > nl[i+1] {
			di = i + 1
			break
		}
	}
	if di == -1 {
		return -1
	}

	for i := 0; i < di; i++ {
		if nl[i] > nl[di] {
			nl[i], nl[di] = nl[di], nl[i]
			break
		}
	}
	reverse(nl[:di])

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
	if len(nl) < 0 {
		return -1
	}

	for i, item := range nl {
		ret += int(math.Pow(10, float64(i))) * int(item)
	}

	//ret = int(nl[len(nl)-1])
	//for i := len(nl) - 2; i >= 0; i-- {
	//	ret = 10*ret + int(nl[i])
	//}

	if ret > int(int32(ret)) {
		return -1
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

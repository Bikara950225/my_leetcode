package main

import "fmt"

func mySqrt(x int) int {
	i := 0
	j := x
	for i < j {
		m := (i + j) >> 1
		if m * m > x {
			j = m
		} else {
			i = m + 1
		}
	}
	if i * i > x {
		i -= 1
	}
	return i
}

func main() {
	ret1 := mySqrt(9)
	if ret1	!= 3 {
		panic(fmt.Errorf(
			"code69 error, not expect result: %d",
			ret1,
		))
	}

	ret2 := mySqrt(8)
	if ret2 != 2 {
		panic(fmt.Errorf(
			"code69 error, not expect result: %d",
			ret2,
		))
	}
}

package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

func baseRand7() int {
	return rand.Intn(8)
}

func baseRand5() int {
	return rand.Intn(5)
}

func rand10() int {
	first := baseRand7()
	for first > 6 {
		first = baseRand7() // 后续会用来判断奇偶，所以概率为二分之一
	}
	second := baseRand7()
	for second > 5 {
		second = baseRand7() // 概率为五分之一
	}

	// (1/2) * (1/5) = 1/10, 符合rand10
	if (first & 0b1) == 0b1 {
		return second + 5
	} else {
		return second
	}
}

func rand100() int {
	first := baseRand7()
	for first >= 4 {
		first = baseRand7()
	}
	second := baseRand7()
	for second >= 5 {
		second = baseRand7()
	}
	trd := baseRand7()
	for trd >= 5 {
		trd = baseRand7()
	}
	return trd + (second * 5) + (first * 25)
}

// 疑似美团考题，用rand5去实现rand7
func rand7() int {
	ret := rand5ToRand10()
	for ret >= 7 {
		ret = rand5ToRand10()
	}
	return ret
}

// 拿 baseRand5 来生成 rand10，给rand7用的
func rand5ToRand10() int {
	first := baseRand5()
	for first > 3 {
		first = baseRand5()
	}
	second := baseRand5()

	if (first & 0b1) == 0b1 {
		return second
	} else {
		return second + 5
	}

	runtime.g
}

func main() {
	for i := 0; i < 10000; i++ {
		r1 := rand7()
		if r1 >= 7 {
			panic(fmt.Errorf("code470 error, result: %d", r1))
		}
	}
}

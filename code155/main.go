package main

import "fmt"

func testCase1(src Stack) {
	src.Push(-2)
	src.Push(0)
	src.Push(-1)

	min1 := src.GetMin()
	if min1 != -2 {
		panic(fmt.Errorf("code155 error, not expect min1: %d\n", min1))
	}

	top1 := src.Top()
	if top1 != -1 {
		panic(fmt.Errorf("code155 error, not expect top1: %d\n", top1))
	}

	src.Pop()

	min2 := src.GetMin()
	if min2 != -2 {
		panic(fmt.Errorf("code155 error, not expect min2: %d\n", min2))

	}
}

func main() {
	testCase1(&MinStack1{})
	testCase1(&MinStack2{})
}

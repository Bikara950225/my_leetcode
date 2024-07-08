package main

func monotoneIncreasingDigits(n int) int {
	var nn []int
	for n > 0 {
		nn = append(nn, n%10)
		n /= 10
	}

	var stack []int
	isLoop := true
	for i := len(nn) - 1; i >= 0; i-- {
		if !isLoop {
			break
		}

		item := nn[i]
		for len(stack) > 0 && item < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			isLoop = false
		}
		stack = append(stack, item)
	}

	for len(stack) < len(nn) {
		stack = append(stack, 9)
	}
	return recoverNum(stack)
}

func recoverNum(stack []int) int {
	ret := stack[0]
	for _, item := range stack[1:] {
		ret = 10*ret + item
	}
	return ret
}

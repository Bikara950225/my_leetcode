package main

func dailyTemperatures(temperatures []int) []int {
	mp := map[int]int{}
	var stack []int
	for i := len(temperatures) - 1; i >= 0; i-- {
		t := temperatures[i]
		for len(stack) > 0 && t >= temperatures[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) <= 0 {
			mp[i] = 0
		} else {
			mp[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	ret := make([]int, len(temperatures))
	for i := range temperatures {
		ret[i] = mp[i]
	}
	return ret
}

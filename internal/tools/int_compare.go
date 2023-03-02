package tools

func IntMax(srcs ...int) int {
	max := srcs[0]
	for i := 1; i < len(srcs); i++ {
		if srcs[i] > max {
			max = srcs[i]
		}
	}
	return max
}

func IntMin(srcs ...int) int {
	min := srcs[0]
	for i := 1; i < len(srcs); i++ {
		if srcs[i] < min {
			min = srcs[i]
		}
	}
	return min
}

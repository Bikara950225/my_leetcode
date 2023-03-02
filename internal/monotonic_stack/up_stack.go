package monotonic_stack

func GetMonoTonicUpStack(src []int) (stack []int) {
	for _, item := range src {
		for len(stack) > 0 && item > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, item)
	}
	return
}

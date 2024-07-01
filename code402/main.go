package main

func removeKdigits(num string, k int) string {
	var stack []byte
	for i := range num {
		n := num[i]

		for len(stack) > 0 && k > 0 && n < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, n)
	}

	for len(stack) > 0 && k > 0 {
		stack = stack[:len(stack)-1]
		k--
	}
	for len(stack) > 0 && stack[0] == '0' {
		stack = stack[1:]
	}
	if len(stack) <= 0 {
		return "0"
	}
	return string(stack)
}

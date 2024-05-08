package main

var leftMap = map[byte]bool{
	'(': true,
	'{': true,
	'[': true,
}

func isValid(s string) bool {
	var stack []byte
	for _, item := range []byte(s) {
		if leftMap[item] {
			stack = append(stack, item)
		} else {
			if len(stack) <= 0 {
				return false
			}

			isPop := false
			switch item {
			case ')':
				isPop = stack[len(stack)-1] == '('
			case '}':
				isPop = stack[len(stack)-1] == '{'
			case ']':
				isPop = stack[len(stack)-1] == '['
			}

			if isPop {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}

	}
	return len(stack) <= 0
}

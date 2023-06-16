package main

func removeDuplicates(s string) string {
	stack := []byte{}
	for _, b := range []byte(s) {
		if len(stack) <= 0 {
			stack = append(stack, b)
		} else {
			tail := len(stack) - 1
			if stack[tail] == b {
				stack = stack[:tail]
			} else {
				stack = append(stack, b)
			}
		}
	}
	return string(stack)
}

func main() {
	ret1 := removeDuplicates("abbaca")
	println(ret1)
}

package main

func smallestSubsequence(s string) string {
	cCount := [26]int{}
	for _, c := range s {
		cCount[c-'a']++
	}

	var stack []byte
	var inStack [26]bool
	for _, c := range s {
		ci := c - 'a'
		cCount[ci]--
		if inStack[ci] {
			continue
		}

		for len(stack) > 0 && c < int32(stack[len(stack)-1]) && cCount[stack[len(stack)-1]-'a'] > 0 {
			tail := len(stack) - 1
			popItem := stack[tail]
			stack = stack[:tail]
			inStack[popItem-'a'] = false
		}
		stack = append(stack, byte(c))
		inStack[c-'a'] = true
	}
	return string(stack)
}

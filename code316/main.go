package main

func removeDuplicateLetters(s string) string {
	bs := [26]int{}
	for i := range s {
		bs[s[i]-'a']++
	}

	var stack []byte
	inStack := [26]bool{}
	for i := range s {
		c := s[i]
		ci := c - 'a'

		bs[ci]--
		if inStack[ci] {
			continue
		}
		for len(stack) > 0 && c < stack[len(stack)-1] && bs[stack[len(stack)-1]-'a'] > 0 {
			inStack[stack[len(stack)-1]-'a'] = false
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, c)
		inStack[ci] = true
	}
	return string(stack)
}

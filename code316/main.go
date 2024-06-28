package main

func removeDuplicateLetters(s string) string {
	// 遍历一次，去重，找到字母最小下标
	cCount := [26]int{}
	for i := 0; i < len(s); i++ {
		cCount[s[i]-'a']++
	}

	// 单调递增
	var stack []byte
	inStack := [26]bool{}
	for i := 0; i < len(s); i++ {
		c := s[i]

		bi := c - 'a'
		// 遍历过要减少次数
		cCount[bi]--

		if inStack[bi] {
			continue
		}
		for len(stack) > 0 && c < stack[len(stack)-1] && cCount[stack[len(stack)-1]-'a'] > 0 {
			// pop
			tail := len(stack) - 1
			popc := stack[tail]
			stack = stack[:tail]
			inStack[popc-'a'] = false
		}
		stack = append(stack, byte(c))
		inStack[c-'a'] = true
	}
	return string(stack)
}

package main

import (
	"fmt"
	"my_leetcode/internal/empty"
)

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	ret := 0
	repeatSet := map[string]empty.Empty{}
	// 滑动窗口
	sIndex := 0
	qIndex := 0
	for sIndex < len(s) && qIndex < len(s) {
		c := string(s[qIndex])
		if _, ok := repeatSet[c]; !ok {
			repeatSet[c] = empty.Empty{}
			if len(repeatSet) > ret {
				ret = len(repeatSet)
			}
			qIndex++
		} else {
			delete(repeatSet, string(s[sIndex]))
			sIndex++
		}
	}
	return ret
}

func lengthOfLongestSubstring2(src string) (ret int) {
	if len(src) <= 0 {
		return
	}
	i, j := 0, 0
	m := map[byte]struct{}{}
	for j < len(src) {
		if _, ok := m[src[j]]; !ok {
			m[src[j]] = struct{}{}
			j++
			if j-i > ret {
				ret = j - i
			}
		} else {
			delete(m, src[i])
			i++
		}
	}

	return
}

func main() {
	ret := lengthOfLongestSubstring2("abcabcbb")
	if ret != 3 {
		panic(fmt.Errorf("%s error", "abcabcbb"))
	}

	ret = lengthOfLongestSubstring2("bbbbb")
	if ret != 1 {
		panic(fmt.Errorf("%s error", "bbbbb"))
	}
}

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

func main() {
	ret := lengthOfLongestSubstring("abcabcbb")
	if ret != 3 {
		panic(fmt.Errorf("%s error", "abcabcbb"))
	}
}

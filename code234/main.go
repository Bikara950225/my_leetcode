package main

import (
	"fmt"
	listnode "my_leetcode/internal/list_node"
)

func main() {
	l := listnode.QuickCreateListNode([]int{1, 2, 2, 1})
	ret1 := isPalindrome1(l)
	if !ret1 {
		panic(fmt.Errorf("code234 isPalindrome() error, not expect result: %t", ret1))
	}
}

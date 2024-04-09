package main

import (
	"fmt"
	treenode "my_leetcode/internal/tree_node"
	"reflect"
)

func main() {
	src := &treenode.TreeNode{
		Val: 10,
		Left: &treenode.TreeNode{
			Val: 5,
			Right: &treenode.TreeNode{
				Val: 6,
			},
		},
		Right: &treenode.TreeNode{
			Val:   12,
			Left:  &treenode.TreeNode{Val: 11},
			Right: &treenode.TreeNode{Val: 13},
		},
	}

	//cc := &CodecPostorder{}
	cc := ConstructorPreorder2()
	srcStr := cc.serialize(src)
	fmt.Println(srcStr)

	deserTree := cc.deserialize(srcStr)
	fmt.Println(deserTree.GetInorder())

	isMatch := reflect.DeepEqual(deserTree, src)
	fmt.Println(isMatch)
}

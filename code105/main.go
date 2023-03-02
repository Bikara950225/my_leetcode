package main

import (
	"fmt"
	treenode "my_leetcode/internal/tree_node"
	"reflect"
)

var (
	inorderMap = map[int]int{}
)

func buildTree(preorder []int, inorder []int) *treenode.TreeNode {
	func() {
		inorderMap = map[int]int{}
	}()

	for i, v := range inorder {
		inorderMap[v] = i
	}
	inorderEnd := len(inorder) - 1
	return dfs(preorder, inorder, 0, inorderEnd, 0, inorderEnd)
}

func dfs(preorder, inorder []int, preBegin, preEnd int, inBegin, inEnd int) *treenode.TreeNode {
	if preBegin > preEnd || inBegin > inEnd {
		return nil
	}

	currVal := preorder[preBegin]
	node := &treenode.TreeNode{
		Val: currVal,
	}
	// 当前指在 inorder 的下标
	currInorderIndex := inorderMap[currVal]
	// 计算left子树的 preorder 的范围
	// (currInorderIndex - 1) - inBegin = leftPreEnd - (preBegin + 1)
	// leftPreEnd = currInorderIndex - inBegin + preBegin
	leftPreEnd := currInorderIndex - inBegin + preBegin

	node.Left = dfs(preorder, inorder, preBegin+1, leftPreEnd, inBegin, currInorderIndex-1)
	node.Right = dfs(preorder, inorder, leftPreEnd+1, preEnd, currInorderIndex+1, inEnd)

	return node
}

func main() {
	preorder1 := []int{3, 9, 20, 15, 7}
	inorder1 := []int{9, 3, 15, 20, 7}

	ret1 := buildTree(preorder1, inorder1)
	ret1List := ret1.GetPreorder()
	if !reflect.DeepEqual(ret1List, []int{3, 9, 20, 15, 7}) {
		panic(fmt.Errorf("code105 error, not expect result: %+v", ret1List))
	}
}

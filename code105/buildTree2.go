package main

import treenode "my_leetcode/internal/tree_node"

func buildTree2(preorder []int, inorder []int) *treenode.TreeNode {
	func() {
		inorderMap = map[int]int{}
	}()

	for idx, item := range inorder {
		inorder[item] = idx
	}

	// preBegin, preEnd; inBegin, inEnd
	return dfs2(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}

func dfs2(preorder []int, inorder []int, preBegin, preEnd int, inBegin, inEnd int) *treenode.TreeNode {
	if preBegin >= preEnd || inBegin >= inEnd {
		return nil
	}

	currVal := preorder[preBegin]
	inIdx := inorderMap[currVal]
	// nextPreEnd - nextPreBegin = nextInEnd - nextInBegin
	nextPreEnd := (inIdx - 1) - inBegin + (preBegin + 1)
	// nextRightPreEnd - nextRightPreBegin = nextRightInEnd - nextRightInBegin
	nextRightPreBegin := preEnd - inEnd + (inIdx + 1)

	return &treenode.TreeNode{
		Val:   currVal,
		Left:  dfs2(preorder, inorder, preBegin+1, nextPreEnd, inBegin, inIdx-1),
		Right: dfs2(preorder, inorder, nextRightPreBegin, preEnd, inIdx+1, inEnd),
	}
}

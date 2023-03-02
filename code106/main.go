package main

import (
	"fmt"
	treenode "my_leetcode/internal/tree_node"
	"reflect"
)

var (
	inorderMap = map[int]int{}
)

func buildTree(inorder []int, postorder []int) *treenode.TreeNode {
	defer func() {
		inorderMap = map[int]int{}
	}()

	// 反转postorder，方便处理而已
	reverse(postorder)

	for i, v := range inorder {
		inorderMap[v] = i
	}

	return dfs(inorder, postorder, 0, len(inorder)-1, 0, len(postorder)-1)
}

func dfs(inorder, postorder []int, inBegin, inEnd int, postBegin, postEnd int) *treenode.TreeNode {
	if inBegin > inEnd || postBegin > postEnd {
		return nil
	}

	currVal := postorder[postBegin]
	currNode := &treenode.TreeNode{Val: currVal}
	currValInorderIndex := inorderMap[currVal]
	// 计算right子树的 postorder 的范围
	// rightPostEnd - (postBegin + 1) = inEnd - (currValInorderIndex + 1)
	// rightPostEnd = inEnd - currValInorderIndex + postBegin
	rightPostEnd := inEnd - currValInorderIndex + postBegin
	currNode.Right = dfs(inorder, postorder, currValInorderIndex+1, inEnd, postBegin+1, rightPostEnd)
	currNode.Left = dfs(inorder, postorder, inBegin, currValInorderIndex-1, rightPostEnd+1, postEnd)

	return currNode
}

func reverse(src []int) {
	i, j := 0, len(src)-1
	for i < j {
		src[i], src[j] = src[j], src[i]
		i++
		j--
	}
}

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	newTree := buildTree(inorder, postorder)
	preorder := newTree.GetPreorder()
	expectPreorder := []int{3, 9, 20, 15, 7}
	if !reflect.DeepEqual(preorder, expectPreorder) {
		panic(fmt.Errorf("code106 error, not expect result: %+v", preorder))
	}
}

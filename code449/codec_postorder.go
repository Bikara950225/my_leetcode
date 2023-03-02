package main

import (
	"fmt"
	"math"
	treenode "my_leetcode/internal/tree_node"
	"strconv"
	"strings"
)

const (
	defaultSplit = " "
)

// 使用后序遍历的方式对二叉搜索树进行序列化、反序列化
type CodecPostorder struct {}

func Constructor() *CodecPostorder {
	return &CodecPostorder{}
}

// Serializes a tree to a single string.
func (s *CodecPostorder) serialize(root *treenode.TreeNode) string {
	// 后序遍历
	preorder := root.GetPostorder()
	preorderStrList := make([]string, 0, len(preorder))
	for i := range preorder {
		preorderStrList = append(preorderStrList, fmt.Sprintf("%d", preorder[i]))
	}
	return strings.Join(preorderStrList, defaultSplit)
}

// Deserializes your encoded data to tree.
func (s *CodecPostorder) deserialize(data string) *treenode.TreeNode {
	data = strings.TrimSpace(data)
	if data == "" {
		return nil
	}

	strList := strings.Split(data, defaultSplit)
	postorder := make([]int, 0, len(strList))
	for _, item := range strList {
		intItem, _ := strconv.Atoi(item)
		postorder = append(postorder, intItem)
	}

	var defFunc func(min, max int) *treenode.TreeNode
	defFunc = func(min, max int) *treenode.TreeNode {
		if len(postorder) <= 0 {
			return nil
		}

		currVal := postorder[len(postorder)-1]
		if currVal < min || currVal > max {
			return nil
		}

		postorder = postorder[:len(postorder)-1]
		return &treenode.TreeNode{
			Val: currVal,
			Right: defFunc(currVal, max),
			Left: defFunc(min, currVal),
		}
	}
	return defFunc(math.MinInt, math.MaxInt)
}

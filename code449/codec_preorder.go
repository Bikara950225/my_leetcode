package main

import (
	"encoding/json"
	"math"
	treenode "my_leetcode/internal/tree_node"
	internalunsafe "my_leetcode/internal/unsafe"
	"strings"
)

// 使用前序遍历的方式对二叉搜索树进行序列化、反序列化
type CodecPreorder struct{}

func ConstructorPreorder() *CodecPreorder {
	return &CodecPreorder{}
}

func (s *CodecPreorder) serialize(root *treenode.TreeNode) string {
	preorderList := root.GetPreorder()
	bs, _ := json.Marshal(preorderList)
	return string(bs)
}

func (s *CodecPreorder) deserialize(data string) *treenode.TreeNode {
	data = strings.TrimSpace(data)
	if data == "" {
		return nil
	}

	var preorderList []int
	err := json.Unmarshal(internalunsafe.StringToByteUnsafe(data), &preorderList)
	if err != nil {
		panic(err)
	}

	// WARN: 这个闭包把 preorderList 当作捕获变量来使用，每次将队头元素出队的时候都会直接影响 preorder
	var dfsFunc func(lower, upper int) *treenode.TreeNode
	dfsFunc = func(lower, upper int) *treenode.TreeNode {
		if len(preorderList) <= 0 {
			return nil
		}
		currVal := preorderList[0]
		if currVal < lower || currVal > upper {
			return nil
		}
		preorderList = preorderList[1:]

		return &treenode.TreeNode{
			Val:   currVal,
			Left:  dfsFunc(lower, currVal),
			Right: dfsFunc(currVal, upper),
		}
	}
	return dfsFunc(math.MinInt, math.MaxInt)
}

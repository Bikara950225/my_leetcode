package main

// 时间复杂度：初始化为 O(1)，其余操作为 O(|S|)，其中 |S|是每次插入或查询的字符串的长度。
// 空间复杂度：O(∣T∣⋅Σ)，其中 ∣T∣ 为所有插入字符串的长度之和，Σ 为字符集的大小，本题Σ=26。
type Trie interface {
	Insert(word string)
	Search(word string) bool
	StartsWith(prefix string) bool
}

type trie struct {
	node [26]*trieNode
}

type trieNode struct {
	sub  [26]*trieNode
	mark bool
}

var _ Trie = (*trie)(nil)

func Constructor() Trie {
	return &trie{}
}

func (s *trie) Insert(word string) {
	if len(word) <= 0 {
		return
	}

	var currNode *trieNode
	nodeList := &(s.node)
	for _, currWord := range word {
		if (*nodeList)[currWord-'a'] == nil {
			(*nodeList)[currWord-'a'] = &trieNode{}
		}
		currNode = (*nodeList)[currWord-'a']
		nodeList = &(currNode.sub)
	}
	if currNode != nil {
		currNode.mark = true
	}
}

func (s *trie) Search(word string) bool {
	nodeList := &(s.node)
	var currNode *trieNode
	for _, currWord := range word {
		currNode = (*nodeList)[currWord-'a']
		if currNode == nil {
			break
		}
		nodeList = &(currNode.sub)
	}
	if currNode != nil {
		return currNode.mark
	}
	return false
}

func (s *trie) StartsWith(prefix string) bool {
	nodeList := &(s.node)
	for _, currWord := range prefix {
		currNode := (*nodeList)[currWord-'a']
		if currNode == nil {
			return false
		}
		nodeList = &(currNode.sub)
	}
	return true
}

package main

type trieAll struct {
	nodeList []*trieAllNode
}

type trieAllNode struct {
	next  []*trieAllNode
	isEnd bool
}

var _ Trie = (*trieAll)(nil)

func NewTrieAll() Trie {
	return &trieAll{
		nodeList: make([]*trieAllNode, 1024),
	}
}

func (s *trieAll) Insert(word string) {
	wordRunes := []rune(word)
	currNodeList := &(s.nodeList)
	var currNode *trieAllNode
	for _, r := range wordRunes {
		if int(r) >= len(*currNodeList) {
			newList := make([]*trieAllNode, r+1)
			copy(newList, *currNodeList)
			*currNodeList = newList
		}
		if (*currNodeList)[r] == nil {
			(*currNodeList)[r] = &trieAllNode{
				next: make([]*trieAllNode, 1024),
			}
		}
		currNode = (*currNodeList)[r]
		currNodeList = &(currNode.next)
	}
	if currNode != nil {
		currNode.isEnd = true
	}
}

func (s *trieAll) Search(word string) bool {
	workRunes := []rune(word)
	currNodeList := s.nodeList
	var ret bool
	for _, r := range workRunes {
		if int(r) > len(currNodeList) {
			ret = false
			break
		}
		if currNodeList[r] == nil {
			ret = false
			break
		}
		ret = currNodeList[r].isEnd
		currNodeList = currNodeList[r].next
	}
	return ret
}

func (s *trieAll) StartsWith(prefix string) bool {
	currNodeList := s.nodeList
	for _, r := range prefix {
		if int(r) > len(currNodeList) {
			return false
		}
		if currNodeList[r] == nil {
			return false
		}
		currNodeList = currNodeList[r].next
	}
	return true
}

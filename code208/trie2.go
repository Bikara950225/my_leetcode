package main

type trie2 struct {
	cache [255]*trie2
	isEnd bool
}

func Constructor2() trie2 {
	return trie2{
		cache: [255]*trie2{},
	}
}

func (this *trie2) Insert(word string) {
	t := this
	for i := 0; i < len(word); i++ {
		r := word[i]
		if t.cache[r] == nil {
			t.cache[r] = &trie2{
				cache: [255]*trie2{},
			}
		}
		t = t.cache[r]
		t.isEnd = i+1 == len(word)
	}
}

func (this *trie2) Search(word string) bool {
	t := this
	for i := 0; i < len(word); i++ {
		r := word[i]
		if t.cache[r] == nil {
			return false
		}
		t = t.cache[r]
	}
	return t.isEnd
}

func (this *trie2) StartsWith(prefix string) bool {
	t := this
	for i := 0; i < len(prefix); i++ {
		r := prefix[i]
		if t.cache[r] == nil {
			return false
		}
		t = t.cache[r]
	}
	return true
}

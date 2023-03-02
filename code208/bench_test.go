package main

import (
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"testing"
)

//go:generate go test --bench=Search$ -benchmem -test.count=3 -test.benchtime=1s .

const (
	strTotal = 10000
)

var (
	srcStringList []string

	srcTrie Trie
	srcMap map[string]struct{}
	isInit bool
)

func init() {
	if isInit {
		return
	}

	srcTrie = NewTrieAll()
	srcMap = map[string]struct{}{}

	for i := 0; i < strTotal; i++ {
		id := uuid.NewV4().String()

		srcStringList = append(srcStringList, id)
		srcTrie.Insert(id)
		srcMap[id] = struct{}{}
	}
	isInit = true
}

func BenchmarkTrie_Search(b *testing.B) {
	randN := rand.Intn(len(srcStringList))

	ok := srcTrie.Search(srcStringList[randN])
	if !ok {
		b.Errorf("TestTrieBench error, not found from trie: %s", srcStringList[randN])
		return
	}
}

func BenchmarkMap_Search(b *testing.B) {
	randN := rand.Intn(len(srcStringList))

	_, ok := srcMap[srcStringList[randN]]
	if !ok {
		b.Errorf("TestTrieBench error, not found from trie: %s", srcStringList[randN])
		return
	}
}

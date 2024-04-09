package sort

import (
	"testing"
)

//go:generate go test --bench=Sort$ -benchmem -count=5 .

var (
	srcList []int
)

func init() {
	srcList = getDescList(100000)
}

func getDescList(n int) []int {
	l := make([]int, 0, n)
	for ; n >= 0; n-- {
		l = append(l, n)
	}
	return l
}

func BenchmarkHeapSort(b *testing.B) {
	l := make([]int, len(srcList))
	copy(l, srcList)

	HeapSort(l)
}

func BenchmarkQuickSort(b *testing.B) {
	l := make([]int, len(srcList))
	copy(l, srcList)

	QuickSort(l)
}

func BenchmarkMergeSort(b *testing.B) {
	l := make([]int, len(srcList))
	copy(l, srcList)

	MergeSort(l)
}

func BenchmarkBubbleSort(b *testing.B) {
	l := make([]int, len(srcList))
	copy(l, srcList)

	BubbleSort(l)
}

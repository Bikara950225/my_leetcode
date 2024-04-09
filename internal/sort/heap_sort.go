package sort

import "my_leetcode/internal/tools"

func HeapSort(src []int) {
	// logN
	// 整体堆排一次，把最大的节点堆到顶
	heapMax2(src)

	// N * logN
	// 把堆顶（当前最大值）放置到列表末端，并且对新的堆顶进行一次堆排（把新的最大值堆顶）
	for i := len(src) - 1; i > 0; {
		tools.IntSwap(&src[0], &src[i])
		i--
		heapMaxisum2(src, 0, i)
	}
}

func heapMax(src []int) {
	for i := len(src) >> 1; i >= 0; i-- {
		heapMaxisum(src, i, len(src)-1)
	}
}

func heapMaxisum(src []int, begin, end int) {
	for (begin<<1)+1 <= end {
		largest := begin

		leftI := (begin << 1) + 1
		if leftI <= end && src[leftI] > src[largest] {
			largest = leftI
		}
		rightI := (begin << 1) + 2
		if rightI <= end && src[rightI] > src[largest] {
			largest = rightI
		}

		if largest != begin {
			tools.IntSwap(&src[begin], &src[largest])
			begin = largest
		} else {
			break
		}
	}
}

func heapMax2(src []int) {
	e := len(src) - 1
	for i := len(src) / 2; i >= 0; i-- {
		heapMaxisum2(src, i, e)
	}
}

func heapMaxisum2(src []int, b, e int) {
	for b<<1+1 <= e {
		largest := b
		left := b<<1 + 1
		if left <= e && src[left] > src[largest] {
			largest = left
		}
		right := b<<1 + 2
		if right <= e && src[right] > src[largest] {
			largest = right
		}

		if largest != b {
			src[b], src[largest] = src[largest], src[b]
			b = largest
		} else {
			break
		}
	}
}

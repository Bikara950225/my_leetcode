package sort

func MergeSort(src []int) {
	mergeSort(src, 0, len(src)-1)
}

func mergeSort(src []int, begin, end int) {
	if begin >= end {
		return
	}

	middle := (end + begin) >> 1
	mergeSort(src, begin, middle)
	mergeSort(src, middle+1, end)

	b1 := begin
	b2 := middle + 1
	tmp := make([]int, 0, end-begin+1)
	for b1 <= middle && b2 <= end {
		if src[b1] < src[b2] {
			tmp = append(tmp, src[b1])
			b1++
		} else {
			tmp = append(tmp, src[b2])
			b2++
		}
	}

	for b1 <= middle {
		tmp = append(tmp, src[b1])
		b1++
	}
	for b2 <= end {
		tmp = append(tmp, src[b2])
		b2++
	}

	for i, item := range tmp {
		src[begin+i] = item
	}
}

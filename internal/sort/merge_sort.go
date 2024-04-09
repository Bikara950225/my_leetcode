package sort

func MergeSort(src []int) {
	mergeSort2(src, 0, len(src)-1)
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

func mergeSort2(nums []int, b, e int) {
	if b >= e {
		return
	}

	m := (b + e) / 2
	mergeSort2(nums, b, m)
	mergeSort2(nums, m+1, e)

	b1, e1 := b, m
	b2, e2 := m+1, e
	tmp := make([]int, 0, e-b+1)
	for b1 <= e1 && b2 <= e2 {
		if nums[b1] <= nums[b2] {
			tmp = append(tmp, nums[b1])
			b1++
		} else {
			tmp = append(tmp, nums[b2])
			b2++
		}
	}
	for b1 <= e1 {
		tmp = append(tmp, nums[b1])
		b1++
	}
	for b2 <= e2 {
		tmp = append(tmp, nums[b2])
		b2++
	}
	for i, v := range tmp {
		nums[b+i] = v
	}
}

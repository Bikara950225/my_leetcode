package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	cpnum1 := make([]int, m)
	copy(cpnum1, nums1)

	i, j := 0, 0
	z := 0
	for i < m && j < n {
		if cpnum1[i] <= nums2[j] {
			nums1[z] = cpnum1[i]
			i++
			z++
		} else if cpnum1[i] > nums2[j] {
			nums1[z] = nums2[j]
			j++
			z++
		}
	}

	for i < m {
		nums1[z] = cpnum1[i]
		i++
		z++
	}
	for j < n {
		nums1[z] = nums2[j]
		j++
		z++
	}
}

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)

	dst1 := make([]int, 6)
	src1 := []int{1, 2, 3, 4, 5, 6}
	copy(dst1, src1)
	fmt.Println(dst1)

	sh := (*reflect.SliceHeader)(unsafe.Pointer(&dst1))
	dst2 := *(*[]int)(unsafe.Pointer(&reflect.SliceHeader{
		Data: sh.Data, Len: 6, Cap: 6,
	}))
	fmt.Println(dst2)
}

package main

import (
	"fmt"
	"reflect"
)

func majorityElement(nums []int) (ret []int) {
	m1, m2 := 0, 0
	v1, v2 := 0, 0

	for _, num := range nums {
		if v1 > 0 && m1 == num {
			v1++
		} else if v2 > 0 && m2 == num {
			v2++
		} else if v1 == 0 {
			m1 = num
			v1++
		} else if v2 == 0 {
			m2 = num
			v2++
		} else {
			v1--
			v2--
		}
	}

	c1, c2 := 0, 0
	for _, num := range nums {
		if v1 > 0 && m1 == num {
			c1++
		}
		if v2 > 0 && m2 == num {
			c2++
		}
	}

	n := len(nums) / 3
	if v1 > 0 && c1 > n {
		ret = append(ret, m1)
	}
	if v2 > 0 && c2 > n {
		ret = append(ret, m2)
	}
	return
}

func main() {
	ret1 := majorityElement([]int{2, 2, 3})
	if !reflect.DeepEqual(ret1, []int{2}) {
		panic(fmt.Errorf("code229 error, ret1 != [2]: %+v", ret1))
	}

	ret2 := majorityElement([]int{2, 3, 2})
	if !reflect.DeepEqual(ret2, []int{2}) {
		panic(fmt.Errorf("code229 error, ret2 != [2]: %+v", ret2))
	}
}

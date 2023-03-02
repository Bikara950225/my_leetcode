package main

import (
	"reflect"
	"testing"
)

func Test_setTop(t *testing.T) {
	t.Run("case1, 设置头节点，不影响尾巴节点", func(t *testing.T) {
		head := &DoubleLinkNode{val: 0}
		var mark *DoubleLinkNode
		for _, val := range [5]int{
			1, 2, 3, 4, 5,
		}{
			newNode := &DoubleLinkNode{val: val}
			head, _ = setTop(head, newNode)
			if val == 3 {
				mark = newNode
			}
		}
		head, _ = setTop(head, mark)
		gotList := head.getListFromNext()
		wantList := []int{3, 5, 4, 2, 1, 0}
		if !reflect.DeepEqual(gotList, wantList) {
			t.Errorf("Test_setTop() error, gotList != wantList: %+v != %+v",
				gotList, wantList,
			)
			return
		}
	})

	t.Run("case2, 设置头节点，影响尾巴节点", func(t *testing.T) {
		head := &DoubleLinkNode{val: 0}
		mark := head
		for _, val := range [5]int{
			1, 2, 3, 4, 5,
		}{
			newNode := &DoubleLinkNode{val: val}
			head, _ = setTop(head, newNode)
		}
		head, tail := setTop(head, mark)
		gotList := head.getListFromNext()
		wantList := []int{0, 5, 4, 3, 2, 1}
		if !reflect.DeepEqual(gotList, wantList) {
			t.Errorf("Test_setTop() error, gotList != wantList: %+v != %+v",
				gotList, wantList,
			)
			return
		}
		if tail.val != 1 {
			t.Errorf("Test_setTop() error, tail is not expect: %+v", *tail)
			return
		}
	})

	t.Run("case3, 头部节点为nil", func(t *testing.T) {
		var head *DoubleLinkNode
		newNode := &DoubleLinkNode{val: 5}

		head, tail := setTop(head, newNode)
		if head != newNode {
			t.Errorf("Test_setTop() error, head isn't expect: %+v", *head)
			return
		}
		if tail != newNode {
			t.Errorf("Test_setTop() error, head isn't expect: %+v", *tail)
			return
		}
	})
}
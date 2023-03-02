package main

import (
	"fmt"
	listnode "my_leetcode/internal/list_node"
	"reflect"
	"testing"
)

func Test_findMiddle(t *testing.T) {
	type args struct {
		head *listnode.ListNode
	}
	tests := []struct {
		name string
		args args
		wantVal int
	}{
		{
			name: "偶数长度的链表",
			args: args{
				head: listnode.QuickCreateListNode([]int{
					1, 2, 3, 4,
				}),
			},
			wantVal: 2,
		},
		{
			name: "奇数长度的链表",
			args: args{
				head: listnode.QuickCreateListNode([]int{
					1, 2, 3, 4, 5,
				}),
			},
			wantVal: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMiddle(tt.args.head)
			if tt.wantVal != got.Val {
				t.Errorf("Test_findMiddle() error, not expect val, %d != %d",
					tt.wantVal, got.Val,
				)
				return
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	t.Run("反转整个链表", func(t *testing.T) {
		link := listnode.QuickCreateListNode([]int{1, 2, 3, 4, 5, 6})
		got := reverse(nil, link, nil)

		wantList := []int{6,5,4,3,2,1}
		gotList := got.GetList()
		if !reflect.DeepEqual(gotList, wantList) {
			t.Errorf("Test_reverse() error, not expect list: %+v", gotList)
			return
		}
	})

	t.Run("反转链表的部分节点", func(t *testing.T) {
		link := listnode.QuickCreateListNode([]int{1, 2, 3, 4, 5, 6})
		pre := link.Next // 2
		curr := pre.Next // 3
		end := curr.Next.Next // 5

		link.Next = reverse(pre, curr, end)
		pre.Next = end

		gotList := link.GetList()
		fmt.Println(gotList)
	})
}
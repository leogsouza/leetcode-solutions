package lc02

import (
	"reflect"
	"testing"
)

type testCase struct {
	l1       []int
	l2       []int
	expected []int
}

func TestAddTwoNumbers(t *testing.T) {
	data := [...]testCase{
		{
			l1:       []int{2, 4, 3},
			l2:       []int{5, 6, 4},
			expected: []int{7, 0, 8},
		},
		{
			l1:       []int{3, 2, 7},
			l2:       []int{4, 3, 2},
			expected: []int{7, 5, 9},
		},
		{
			l1:       []int{6, 9, 5},
			l2:       []int{5, 6, 6},
			expected: []int{1, 6, 2, 1},
		},
	}

	for _, tt := range data {
		l1 := sliceToListNode(tt.l1)
		l2 := sliceToListNode(tt.l2)
		got := listNodeToSlice(addTwoNumber(l1, l2))
		if !reflect.DeepEqual(got, tt.expected) {
			t.Fatalf("in: %v %v, got: %v, expected; %v", tt.l1, tt.l2, got, tt.expected)
		}
	}
}

func listNodeToSlice(l *ListNode) (s []int) {
	for l != nil {
		s = append(s, l.Val)
		l = l.Next
	}
	return
}

func sliceToListNode(s []int) *ListNode {
	l := &ListNode{}
	t := l
	for _, v := range s {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}

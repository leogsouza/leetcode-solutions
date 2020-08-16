package addtwonumbers

// ListNode is a definition for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumber(l1, l2 *ListNode) *ListNode {
	result, aux := &ListNode{}, 0
	temp := result
	for l1 != nil || l2 != nil || aux != 0 {
		if l1 != nil {
			aux += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			aux += l2.Val
			l2 = l2.Next
		}
		temp.Next = &ListNode{Val: aux % 10}
		temp = temp.Next
		aux /= 10
	}

	return result.Next
}

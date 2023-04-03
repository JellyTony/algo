package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

// addTwoNumbers 两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	list := &ListNode{}
	result := list
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		list.Next = &ListNode{Val: carry % 10}
		carry = carry / 10
		list = list.Next
	}

	return result.Next
}

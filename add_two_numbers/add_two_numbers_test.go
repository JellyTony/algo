package add_two_numbers

import (
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	node := addTwoNumbers(
		&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
		&ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
	)

	for node != nil {
		t.Log(node.Val)
		node = node.Next
	}
}

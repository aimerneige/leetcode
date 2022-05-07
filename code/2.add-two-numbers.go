package code

/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

// ListNode list node
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := false
	var head *ListNode
	p := head
	for !(l1 == nil && l2 == nil && !carry) {
		sum := 0
		if carry {
			sum++
		}
		if l1 != nil {
			sum += l1.Val
		}
		if l2 != nil {
			sum += l2.Val
		}
		if sum > 9 {
			carry = true
			sum = sum % 10
		} else {
			carry = false
		}
		node := &ListNode{
			Val:  sum,
			Next: nil,
		}
		if head == nil {
			head = node
			p = node
		} else {
			p.Next = node
			p = p.Next
		}

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return head
}

// @lc code=end

package code

/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 */

// @lc code=start
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head
	pre := p
	p = p.Next
	for p != nil {
		if p.Val == pre.Val {
			pre.Next = p.Next
			p = p.Next
			continue
		}
		pre = pre.Next
		p = p.Next
	}

	return head
}

// @lc code=end

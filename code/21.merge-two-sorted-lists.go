package code

/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

// @lc code=start
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var head *ListNode
	var p *ListNode
	for {
		if list1 == nil {
			p.Next = list2
			break
		}
		if list2 == nil {
			p.Next = list1
			break
		}
		v1 := list1.Val
		v2 := list2.Val
		if v1 <= v2 {
			if head == nil {
				head = list1
				p = head
				list1 = list1.Next
				p.Next = nil
			} else {
				p.Next = list1
				list1 = list1.Next
				p = p.Next
				p.Next = nil
			}
		} else {
			if head == nil {
				head = list2
				p = head
				list2 = list2.Next
				p.Next = nil
			} else {
				p.Next = list2
				list2 = list2.Next
				p = p.Next
				p.Next = nil
			}
		}
	}
	return head
}

// @lc code=end

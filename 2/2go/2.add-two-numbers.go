/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	current := head
	sums := 0

	for l1 != nil || l2 != nil || sums > 9{
		sums /= 10
		if l1 != nil {
			sums += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sums += l2.Val
			l2 = l2.Next
		}

		if sums > 9 {
			sums = sums - 10
			current.Next = &ListNode{Val: sums}
			sums = sums + 10
		}else{
			current.Next = &ListNode{Val: sums}
		}

		current = current.Next
	}	
	return head.Next
}
// @lc code=end


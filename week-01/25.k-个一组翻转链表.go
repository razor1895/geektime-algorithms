/*
 @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
*/

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{0, head}
	prev := dummy
	end := dummy

	for end != nil {
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}
		start := prev.Next
		next := end.Next
		end.Next = nil
		prev.Next = reverse(start)
		start.Next = next
		prev = start

		end = prev
	}

	return dummy.Next
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode = nil
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=515 lang=golang
 *
 * [515] 在每个树行中找最大值
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	q := []*TreeNode{root}

	for len(q) > 0 {
		size := len(q)
		v := q[0].Val
		for size > 0 {
			node := q[0]
			q = q[1:]
			v = max(v, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			size--
		}

		res = append(res, v)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

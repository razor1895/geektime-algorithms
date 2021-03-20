/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N 叉树的层序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	res := make([][]int, 0)

	if root == nil {
		return res
	}

	q := []*Node{root}

	for len(q) > 0 {
		level := []int{}
		qLen := len(q)
		for i := 0; i < qLen; i++ {
			node := q[0]
			q = q[1:]
			level = append(level, node.Val)
			if node.Children != nil {
				q = append(q, node.Children...)
			}
		}
		res = append(res, level)
	}

	return res
}

// @lc code=end

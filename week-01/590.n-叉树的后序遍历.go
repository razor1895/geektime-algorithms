/*
 * @lc app=leetcode.cn id=590 lang=golang
 *
 * [590] N 叉树的后序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
func postorder(root *Node) []int {
	res := []int{}
	if root == nil {
		return res
	}

	stack := []*Node{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append([]int{node.Val}, res...)

		for _, v := range node.Children {
			stack = append(stack, v)
		}
	}

	return res
}

// @lc code=end

// package main

/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N 叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
//  * type Node struct {
 *     Val int
 *     Children []*Node
 * }
*/

// type Node struct {
// 	Val      int
// 	Children []*Node
// }

func preorder(root *Node) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	stack := []*Node{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}

	return res
}

// @lc code=end

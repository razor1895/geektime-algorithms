/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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
func isValidBST(root *TreeNode) bool {
	pre := math.MinInt64
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		if !dfs(node.Left) {
			return false
		}
		if node.Val <= pre {
			return false
		}
		pre = node.Val

		return dfs(node.Right)
	}

	return dfs(root)
}

// @lc code=end

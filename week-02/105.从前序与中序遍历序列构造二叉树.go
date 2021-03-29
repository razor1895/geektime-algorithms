package main

/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

var hash map[int]int

func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	hash = make(map[int]int)

	for i, v := range inorder {
		hash[v] = i
	}

	return buildHelper(preorder, inorder, 0, n-1, 0, n-1)
}

func buildHelper(preorder, inorder []int, preStart, preEnd, inStart, inEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}

	root := &TreeNode{Val: preorder[preStart]}
	rootIndex := hash[root.Val]
	leftSize := rootIndex - inStart
	root.Left = buildHelper(preorder, inorder, preStart+1, preStart+leftSize, inStart, rootIndex-1)
	root.Right = buildHelper(preorder, inorder, preStart+leftSize+1, preEnd, rootIndex+1, inEnd)
	return root
}

// @lc code=end

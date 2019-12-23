package leetcode

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := float64(maxDepth(root.Left))
	right := float64(maxDepth(root.Right))
	return 1 + int(math.Max(left, right))
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := float64(minDepth(root.Left))
	right := float64(minDepth(root.Right))

	if root.Left != nil && root.Right != nil {
		return 1 + int(math.Min(left, right))
	} else {
		return 1 + int(left) + int(right)
	}
}

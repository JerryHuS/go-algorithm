/**
 * @Author: alessonhu
 * @Description:
 * @File:  最长节点距离
 * @Version: 1.0.0
 * @Date: 2022/4/25 15:40
 */

package main

func diameterOfBinaryTree(root *TreeNode) int {
	maxDia := 0
	if root == nil {
		return maxDia
	}
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		lh := dfs(node.Left)
		rh := dfs(node.Right)
		maxDia = max(maxDia, lh+rh)
		return 1 + max(lh, rh)
	}
	dfs(root)
	return maxDia
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

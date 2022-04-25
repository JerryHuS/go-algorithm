/**
 * @Author: alessonhu
 * @Description:
 * @File:  是否对称
 * @Version: 1.0.0
 * @Date: 2022/4/25 15:38
 */

package main

func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}

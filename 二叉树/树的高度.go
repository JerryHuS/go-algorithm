/**
 * @Author: alessonhu
 * @Description:
 * @File:  树的高度
 * @Version: 1.0.0
 * @Date: 2022/4/26 10:42
 */
package main

import "container/list"

func getTreeHeightDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return maxInt(getTreeHeightDFS(root.Left), getTreeHeightDFS(root.Right)) + 1
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getTreeHeightBFS(root *TreeNode) int {
	ans := 0
	if root == nil {
		return ans
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		lenTmp := queue.Len()
		for lenTmp > 0 {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			lenTmp--
		}
		ans++
	}
	return ans
}

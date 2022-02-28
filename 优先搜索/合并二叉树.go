package 优先搜索

import (
	"container/list"
	"fmt"
)

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

//BFS
func mergeTreesBFS(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	q1 := list.New()
	q1.PushBack(root1)
	q2 := list.New()
	q2.PushBack(root2)

	mergeNode := &TreeNode{
		Val:   root1.Val + root2.Val,
		Left:  nil,
		Right: nil,
	}
	mergeQ := list.New()
	mergeQ.PushBack(mergeNode)
	for q1.Len() > 0 && q2.Len() > 0 {
		node := mergeQ.Front().Value.(*TreeNode)
		mergeQ.Remove(mergeQ.Front())
		node1 := q1.Front().Value.(*TreeNode)
		q1.Remove(q1.Front())
		node2 := q2.Front().Value.(*TreeNode)
		q2.Remove(q2.Front())
		left1, right1 := node1.Left, node1.Right
		left2, right2 := node2.Left, node2.Right
		if left1 != nil || left2 != nil {
			if left1 != nil && left2 != nil {
				left := &TreeNode{
					Val:   left1.Val + left2.Val,
					Left:  nil,
					Right: nil,
				}
				node.Left = left
				mergeQ.PushBack(left)
				q1.PushBack(left1)
				q2.PushBack(left2)
			} else if left1 != nil {
				node.Left = left1
			} else {
				node.Left = left2
			}
		}

		if right1 != nil || right2 != nil {
			if right1 != nil && right2 != nil {
				right := &TreeNode{
					Val:   right1.Val + right2.Val,
					Left:  nil,
					Right: nil,
				}
				node.Right = right
				mergeQ.PushBack(right)
				q1.PushBack(right1)
				q2.PushBack(right2)
			} else if right1 != nil {
				node.Right = right1
			} else {
				node.Right = right2
			}
		}
	}
	return mergeNode
}

//DFS
func mergeTreesDFS(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val = root1.Val + root2.Val
	root1.Left = mergeTreesBFS(root1.Left, root2.Left)
	root1.Right = mergeTreesBFS(root1.Right, root2.Right)
	return root1
}

func BFSTrees1(root *TreeNode) []int {
	valArr := make([]int, 0)
	q := list.New()
	q.PushBack(root)
	for q.Len() > 0 {
		for i := 0; i < q.Len(); i++ {
			node := q.Remove(q.Front()).(*TreeNode)
			if node == nil {
				continue
			}
			fmt.Println(node)
			valArr = append(valArr, node.Val)
			q.PushBack(node.Left)
			q.PushBack(node.Right)
		}
	}
	return valArr
}

func BFSTrees(root *TreeNode) []int {
	valArr := make([]int, 0)
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		valArr = append(valArr, node.Val)
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return valArr
}

func DFSTrees(root *TreeNode) []int {
	valArr := make([]int, 0)
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		valArr = append(valArr, node.Val)
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
	return valArr
}

func BFSZTrees(root *TreeNode) []int {
	arrInt := make([]int, 0)
	queue := list.New()
	queue.PushBack(root)
	i := 0
	for queue.Len() > 0 {
		tmpQueue := queue
		i++
		queue = list.New()
		for tmpQueue.Len() > 0 {
			node := tmpQueue.Remove(tmpQueue.Front()).(*TreeNode)
			arrInt = append(arrInt, node.Val)
			if i%2 != 0 {
				if node.Left != nil {
					queue.PushBack(node.Left)
				}
				if node.Right != nil {
					queue.PushBack(node.Right)
				}
			} else {
				if node.Right != nil {
					queue.PushBack(node.Right)
				}
				if node.Left != nil {
					queue.PushBack(node.Left)
				}
			}
		}
	}
	return arrInt
}

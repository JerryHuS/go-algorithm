package 优先搜索

import "container/list"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 层次遍历
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		tmp := queue
		queue = list.New()
		for tmp.Len() > 0 {
			node := tmp.Front().Value.(*Node)
			tmp.Remove(tmp.Front())
			if tmp.Len() > 0 {
				node.Next = tmp.Front().Value.(*Node)
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return root
}

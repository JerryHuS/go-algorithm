/**
 * @Author: alessonhu
 * @Description:
 * @File:  sortList
 * @Version: 1.0.0
 * @Date: 2022/4/18 15:56
 */

package main

import "sort"

func sortList(root *ListNode) *ListNode {
	arr := make([]int, 0)
	for root != nil {
		arr = append(arr, root.Val)
		root = root.Next
	}
	sort.Ints(arr)

	var head, cur *ListNode
	for i, _ := range arr {
		if head == nil {
			head = &ListNode{Val: arr[i]}
			cur = head
		} else {
			cur.Next = &ListNode{Val: arr[i]}
			cur = cur.Next
		}
	}
	return head
}

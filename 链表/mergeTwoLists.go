/**
 * @Author: alessonhu
 * @Description:
 * @File:  mergeTwoLists
 * @Version: 1.0.0
 * @Date: 2022/4/18 14:33
 */

package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode
	var cur *ListNode
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			if head == nil {
				head = &ListNode{Val: list1.Val}
				cur = head
			} else {
				cur.Next = &ListNode{Val: list1.Val}
				cur = cur.Next
			}
			list1 = list1.Next
		} else {
			if head == nil {
				head = &ListNode{Val: list2.Val}
				cur = head
			} else {
				cur.Next = &ListNode{Val: list2.Val}
				cur = cur.Next
			}
			list2 = list2.Next
		}
	}
	for list1 != nil {
		if head == nil {
			head = &ListNode{Val: list1.Val}
			cur = head
		} else {
			cur.Next = &ListNode{Val: list1.Val}
			cur = cur.Next
		}
		list1 = list1.Next
	}
	for list2 != nil {
		if head == nil {
			head = &ListNode{Val: list2.Val}
			cur = head
		} else {
			cur.Next = &ListNode{Val: list2.Val}
			cur = cur.Next
		}
	}
	return head
}

func mergeTwoListsFast(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for list1 != nil || list2 != nil {
		if list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				cur.Next = list1
				list1 = list1.Next
			} else {
				cur.Next = list2
				list2 = list2.Next
			}
			cur = cur.Next
		} else if list1 != nil {
			cur.Next = list1
			break
		} else {
			cur.Next = list2
			break
		}
	}
	return head.Next
}

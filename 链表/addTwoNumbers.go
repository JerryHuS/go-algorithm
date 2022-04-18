/**
 * @Author: alessonhu
 * @Description:
 * @File:  addTwoNumbers
 * @Version: 1.0.0
 * @Date: 2022/4/18 11:27
 */

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, cur *ListNode
	plus := 0
	for l1 != nil || l2 != nil {
		a, b := 0, 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		val := (a + b + plus) % 10
		plus = (a + b + plus) / 10
		if head == nil {
			head = &ListNode{Val: val}
			cur = head
		} else {
			cur.Next = &ListNode{Val: val}
			cur = cur.Next
		}
	}
	if plus > 0 {
		cur.Next = &ListNode{Val: plus}
	}
	return head
}

/**
 * @Author: alessonhu
 * @Description:
 * @File:  removeNthFromEnd
 * @Version: 1.0.0
 * @Date: 2022/4/18 14:24
 */

package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	len := 0
	tmp := head
	for tmp != nil {
		len++
		tmp = tmp.Next
	}

	if n == len {
		return head.Next
	}

	delPre := head
	for i := 1; i < len-n; i++ {
		delPre = delPre.Next
	}

	delPre.Next = delPre.Next.Next

	return head
}

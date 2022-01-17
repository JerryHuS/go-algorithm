package 双指针

import (
	"sort"
)

/*
示例 1：

输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]
排序后，数组变为 [0,1,9,16,100]

示例 2：
输入：nums = [-7,-3,2,3,11]
输出：[4,9,9,49,121]
*/

//1、有序数组的平方 O（nlogn）
func sortedSquares1(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
	}
	sort.Ints(nums)
	return nums
}

// 输入是升序，找正负数分界线 O（n）
func sortedSquares2(nums []int) []int {
	return nums
}

// 双指针 O（n）
func sortedSquares3(nums []int) []int {
	n := len(nums)
	i, j := 0, n-1
	ans := make([]int, n)
	for pos := n - 1; pos >= 0; pos-- {
		if m, n := nums[i]*nums[i], nums[j]*nums[j]; m < n {
			ans[pos] = n
			j--
		} else {
			ans[pos] = m
			i++
		}
	}
	return ans
}

//2、给你一个数组，将数组中的元素向右轮转 k 个位置，其中 k 是非负数
func rotate(nums []int, k int) {
	n := len(nums)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[(i+k)%n] = nums[i]
	}
	copy(nums, ans)
}

func rotate1(nums []int, k int) {
	n := len(nums)
	k %= n
	for start, count := 0, gcd(k, n); start < count; start++ {
		pre, cur := nums[start], start
		for ok := true; ok; ok = cur != start {
			next := (cur + k) % n
			nums[next], pre, cur = pre, nums[next], next
		}
	}
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

//3、移动零，将数组中的0移到末尾
// 0 0 1
func moveZeroes(nums []int) {
	for i, j := 0, 0; i < len(nums); {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
		i++
	}
}

//*
// * Definition for singly-linked list.
// * type ListNode struct {
// *     Val int
// *     Next *ListNode
// * }
//

type ListNode struct {
	Val  int
	Next *ListNode
}

//4、输出链表的中间节点，如果有两个就输出第二个
//12345 3
//1234 3
//123 2
func middleNode1(head *ListNode) *ListNode {
	lenTmp := head
	len := 0
	for lenTmp != nil {
		len++
		lenTmp = lenTmp.Next
	}
	midIndex := len/2 + 1
	for i := 1; i < midIndex; i++ {
		head = head.Next
	}
	return head
}

//12345 3
//1234 3
//123 2
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

//5、删除链表的倒数第 N 个结点
// 1 2 3 4 5
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	len := 0
	tmp := head
	for tmp != nil {
		len++
		tmp = tmp.Next
	}
	//12345 2
	//12 2
	//123 1
	i := len - n + 1

	dummy := &ListNode{0, head}
	dummyTmp := dummy

	for j := 1; j < i; j++ {
		dummyTmp = dummyTmp.Next
	}

	if dummyTmp.Next.Next == nil {
		dummyTmp.Next = nil
	} else {
		dummyTmp.Next = dummyTmp.Next.Next
	}
	return dummy.Next
}

//快慢指针
//12345  2   n+m=L
//  ｜
//012345
//
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := head, dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ; first != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}

/*
总结：
双向指针:
-收尾指针
-快慢指针
*/

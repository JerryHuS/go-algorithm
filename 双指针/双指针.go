package 双指针

import "sort"

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

package 双指针

import (
	"sort"
)

/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，
使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]

输入：nums = []
输出：[]

输入：nums = [0]
输出：[]

*/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	len := len(nums)
	for x := 0; x < len; x++ {
		if !(x == 0 || nums[x] != nums[x-1]) {
			continue
		}
		end := len - 1
		ans := -nums[x]
		for y := x + 1; y < len; y++ {
			if !(y == x+1 || nums[y] != nums[y-1]) {
				continue
			}
			for y < end && nums[y]+nums[end] > ans {
				end--
			}
			if y == end {
				break
			}
			if nums[y]+nums[end] == ans {
				result = append(result, []int{nums[y], nums[end], nums[x]})
			}
		}
	}
	return result
}

/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
*/
func twoSum(nums []int, target int) []int {
	tmpMap := make(map[int]int)
	for i, _ := range nums {
		if j, ok := tmpMap[target-nums[i]]; ok {
			return []int{i, j}
		}
		tmpMap[nums[i]] = i
	}
	return nil
}

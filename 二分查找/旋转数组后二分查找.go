/**
 * @Author: alessonhu
 * @Description:
 * @File:  旋转数组后二分查找
 * @Version: 1.0.0
 * @Date: 2022/4/26 10:42
 */
package main

import "math"

/*
1 2 3 4 5 6 7
[]int{2,3,4,5,6,7,1}
[]int{6,7,1,2,3,4,5}
*/
func findTarget(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] == target {
			return middle
		}
		if nums[middle] > nums[left] {
			if nums[left] < target && target < nums[middle] {
				right = middle - 1
			} else {
				left = middle + 1
			}
		} else {
			if nums[middle] < target && target < nums[right] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		}
	}
	return -1
}

func findMoveK(nums []int) int {
	min := math.MaxInt32
	ans := 0
	for i := range nums {
		if nums[i] < min {
			min = nums[i]
			ans = i
		}
	}
	return ans
}

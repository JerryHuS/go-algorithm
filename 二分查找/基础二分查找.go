package main

//1、nums升序，返回命中的下标值，不存在返回-1
//-1, 0, 1, 2, 3, 4
func search(nums []int, target int) int {
	i, j := 0, len(nums)-1
	mid := i
	for i <= j {
		mid = (j-i)/2 + i
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return -1
}

/**
	2、判断第一个坏版本
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
*/

func isBadVersion(version int) bool {
	return false
}

func firstBadVersion(n int) int {
	i, j := 1, n
	mid := i
	for i < j {
		mid = (j-i)/2 + i
		if isBadVersion(mid) {
			j = mid
		} else {
			i = mid + 1
		}
	}
	//sort.Search(n, func(i int) bool {
	//	return isBadVersion(i)
	//})
	return i
}

//3、搜索插入位置
func searchInsert(nums []int, target int) int {
	i, j := 0, len(nums)-1
	mid := i
	for i < j {
		mid = (j-i)/2 + i
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	if nums[i] < target {
		return i + 1
	} else {
		return i
	}
}

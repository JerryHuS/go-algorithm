package 双指针

import "math"

//覆盖的最大面积
func maxArea(height []int) int {
	max := 0
	begin, end := 0, len(height)-1
	for begin < end {
		if size := (end - begin) * int(math.Min(float64(height[begin]), float64(height[end]))); size > max {
			max = size
		}
		if height[begin] > height[end] {
			end--
		} else {
			begin++
		}
	}
	return max
}

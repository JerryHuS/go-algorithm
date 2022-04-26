/**
 * @Author: alessonhu
 * @Description:
 * @File:  旋转数组后二分查找_test.go
 * @Version: 1.0.0
 * @Date: 2022/4/26 10:58
 */
package main

import "testing"

func Test_findTarget(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "case1", args: args{[]int{2, 3, 4, 5, 6, 7, 1}, 1}, want: 6},
		{name: "case2", args: args{[]int{2, 3, 4, 5, 6, 7, 1}, 10}, want: -1},
		{name: "case3", args: args{[]int{6, 7, 1, 2, 3, 4, 5}, 7}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTarget(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("findTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}

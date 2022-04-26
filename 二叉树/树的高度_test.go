/**
 * @Author: alessonhu
 * @Description:
 * @File:  树的高度_test.go
 * @Version: 1.0.0
 * @Date: 2022/4/26 15:01
 */
package main

import "testing"

func Test_getTreeHeightBFS(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// [1,2,3,4,null,null,5]
		/*
			            1
					2       3
			    4               5
		*/
		{name: "case1", args: args{root: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &TreeNode{
				Val:  3,
				Left: nil,
				Right: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
			},
		}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTreeHeightBFS(tt.args.root); got != tt.want {
				t.Errorf("getTreeHeightBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

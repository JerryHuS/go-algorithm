package 优先搜索

import (
	"reflect"
	"testing"
)

func TestBFSTrees(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "case1", args: args{root: &TreeNode{
			Val:   1,
			Left:  &TreeNode{99, nil, nil},
			Right: &TreeNode{100, nil, &TreeNode{4, nil, nil}},
		}}, want: []int{1, 99, 100, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BFSTrees(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BFSTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDFSTrees(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "case1", args: args{root: &TreeNode{
			Val:   1,
			Left:  &TreeNode{99, &TreeNode{4, nil, nil}, nil},
			Right: &TreeNode{100, nil, nil},
		}}, want: []int{1, 99, 4, 100}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DFSTrees(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DFSTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeTreesBFS(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{name: "case1", args: args{
			/*
				        1
					99		100
				4

						1
					99		100
						4
			*/
			root1: &TreeNode{
				Val:   1,
				Left:  &TreeNode{99, &TreeNode{4, nil, nil}, nil},
				Right: &TreeNode{100, nil, nil},
			},
			root2: &TreeNode{
				Val:   1,
				Left:  &TreeNode{99, nil, &TreeNode{4, nil, nil}},
				Right: &TreeNode{100, nil, nil},
			}},
			want: &TreeNode{
				Val:   2,
				Left:  &TreeNode{198, &TreeNode{4, nil, nil}, &TreeNode{4, nil, nil}},
				Right: &TreeNode{200, nil, nil},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTreesBFS(tt.args.root1, tt.args.root2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTreesBFS() = %v, want %v", got, tt.want)
			}
			if got := mergeTreesDFS(tt.args.root1, tt.args.root2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTreesBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

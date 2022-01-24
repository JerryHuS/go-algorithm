package 优先搜索

import (
	"reflect"
	"testing"
)

func Test_updateMatrix(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		//输入：mat = [[0,0,0],[0,1,0],[1,1,1]]
		//输出：		 [[0,0,0],[0,1,0],[1,2,1]]
		{name: "case1", args: args{mat: [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}}, want: [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}}},
		//[[0],[0],[0],[0],[0]]
		{name: "case2", args: args{mat: [][]int{{0}, {0}, {0}, {0}, {0}}}, want: [][]int{{0}, {0}, {0}, {0}, {0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateMatrix(tt.args.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

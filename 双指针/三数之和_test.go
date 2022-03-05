package 双指针

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{name: "case1", args: args{nums: []int{-1, 0, 1, 2, -1, -4}}, want: [][]int{{-1, 2, -1}, {-1, 0, 1}}},
		{name: "case2", args: args{nums: []int{}}, want: [][]int{}},
		{name: "case3", args: args{nums: []int{0}}, want: [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

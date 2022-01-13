package 双指针

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_sortedSquares(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "case1", args: args{nums: []int{-2, 1, 2, 3}}, want: []int{1, 4, 4, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares1(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortedSquares3(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "case1", args: args{nums: []int{-2, 1, 2, 3}}, want: []int{1, 4, 4, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares3(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "case1", args: args{nums: []int{1, 2, 3, 4}, k: 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("origin:", tt.args.nums)
			rotate(tt.args.nums, tt.args.k)
			fmt.Println("rotate:", tt.args.nums)
		})
	}
}

package 二分查找

import "testing"

func Test_search(t *testing.T) {
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
		{name: "case1", args: args{nums: []int{-1, 0, 1, 2, 3, 4}, target: 0}, want: 1},
		{name: "case2", args: args{nums: []int{-1, 99, 100, 200}, target: 100}, want: 2},
		{name: "case3", args: args{nums: []int{-1, 0, 1, 2, 3, 4}, target: 10}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchInsert(t *testing.T) {
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
		{name: "case1", args: args{nums: []int{1, 2, 3, 4}, target: 0}, want: 0},
		{name: "case2", args: args{nums: []int{1, 2, 3, 4}, target: 2}, want: 1},
		{name: "case3", args: args{nums: []int{1, 2, 3, 4}, target: 10}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

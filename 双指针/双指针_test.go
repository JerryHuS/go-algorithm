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

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "case1", args: args{nums: []int{0, 1, 2, 3}}},
		{name: "case2", args: args{nums: []int{0, 1, 2, 3, 0, 0}}},
		{name: "case3", args: args{nums: []int{1, 2, 3, 0, 0}}},
		{name: "case4", args: args{nums: []int{0, 0, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.args.nums)
			moveZeroes(tt.args.nums)
			fmt.Println(tt.args.nums)
		})
	}
}

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
				n: 1,
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
		},
		{
			name: "case2",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
				n: 2,
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
		{
			name: "case3",
			args: args{
				head: &ListNode{
					Val:  1,
					Next: nil,
				},
				n: 1,
			},
			want: nil,
		},
		{
			name: "case4",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: nil,
					},
				},
				n: 2,
			},
			want: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %+v, want %v", got, tt.want)
				fmt.Print("got:")
				PrintList(got)
			}
		})
	}
}

func PrintList(list *ListNode) {
	for list != nil {
		fmt.Print(list.Val)
		list = list.Next
	}
	fmt.Println()
}

func Test_middleNode(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "case1", args: args{head: &ListNode{1, &ListNode{2, &ListNode{3, nil}}}}, want: &ListNode{2, &ListNode{3, nil}},
		},
		{
			name: "case2", args: args{head: &ListNode{1, &ListNode{2, nil}}}, want: &ListNode{2, nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

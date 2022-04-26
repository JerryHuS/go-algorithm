package main

import (
	"fmt"
	"testing"
)

func Test_sortByBubble(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "case1", args: args{[]int{1, 3, 2, 5, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortByBubble(tt.args.arr)
			fmt.Println(tt.args.arr)
		})
	}
}

func Test_sortByQuick(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "case1", args: args{[]int{1, 3, 2, 5, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortByQuick(tt.args.arr)
			fmt.Println(tt.args.arr)
		})
	}
}

package 优先搜索

import "testing"

func Test_maxAreaOfIsland(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "case1", args: args{grid: [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}}, want: 6},
		{name: "case2", args: args{grid: [][]int{{0, 0, 0}, {1, 1, 0}, {1, 0, 1}}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAreaOfIsland(tt.args.grid); got != tt.want {
				t.Errorf("maxAreaOfIsland() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxAreaOfIslandBFS(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "case1", args: args{grid: [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}}, want: 6},
		{name: "case2", args: args{grid: [][]int{{0, 0, 0}, {1, 1, 0}, {1, 0, 1}}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAreaOfIslandBFS(tt.args.grid); got != tt.want {
				t.Errorf("maxAreaOfIslandBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

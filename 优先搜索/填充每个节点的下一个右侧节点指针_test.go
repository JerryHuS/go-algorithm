package 优先搜索

import (
	"reflect"
	"testing"
)

func Test_connect(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		// TODO: Add test cases.
		{name: "case1", args: args{root: &Node{
			Val: 1,
			Left: &Node{
				Val:   2,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Right: &Node{
				Val:   3,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Next: nil,
		}},
			want: &Node{
				Val: 1,
				Left: &Node{
					Val:   2,
					Left:  nil,
					Right: nil,
					Next:  nil,
				},
				Right: &Node{
					Val:   3,
					Left:  nil,
					Right: nil,
					Next:  nil,
				},
				Next: nil,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connect(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connect() = %v, want %v", got, tt.want)
			}
		})
	}
}

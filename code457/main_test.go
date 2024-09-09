package main

import "testing"

func Test_circularArrayLoop(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "path1",
			args: args{
				nums: []int{2, -1, 1, 2, 2},
			},
			want: true,
		}, {
			name: "path2",
			args: args{
				nums: []int{-2, 1, -1, -2, -2},
			},
			want: false,
		}, {
			name: "path3",
			args: args{
				nums: []int{-1, -2, -3, -4, -5, 6},
			},
			want: false,
		},
		{
			name: "path4",
			args: args{
				nums: []int{1, -1, 5, 1, 4},
			},
			want: true,
		},
		{
			name: "path5",
			args: args{
				nums: []int{-1},
			},
			want: false,
		},
		{
			name: "path6",
			args: args{
				nums: []int{-1, -1, -3},
			},
			want: false,
		},
		{
			name: "path7",
			args: args{
				nums: []int{-1},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := circularArrayLoop(tt.args.nums); got != tt.want {
				t.Errorf("circularArrayLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}

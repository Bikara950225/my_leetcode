package tools

import "testing"

func TestIntMax(t *testing.T) {
	type args struct {
		srcs []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "happy path1",
			args: args{
				srcs: []int{1, 2, 3, 4, 5},
			},
			want: 5,
		},
		{
			name: "happy path2",
			args: args{
				srcs: []int{1, 1, 1, 1, 1, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntMax(tt.args.srcs...); got != tt.want {
				t.Errorf("IntMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntMin(t *testing.T) {
	type args struct {
		srcs []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "happy path1",
			args: args{
				srcs: []int{5, 4, 3, 1, 2},
			},
			want: 1,
		},
		{
			name: "happy path2",
			args: args{
				srcs: []int{2, 2, 2, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntMin(tt.args.srcs...); got != tt.want {
				t.Errorf("IntMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

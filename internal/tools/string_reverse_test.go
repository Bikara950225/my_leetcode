package tools

import "testing"

func TestStringReverse(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "happy path1, 英文",
			args: args{
				src: "hello",
			},
			want: "olleh",
		},
		{
			name: "happy path1, 中文",
			args: args{
				src: "你好 测试",
			},
			want: "试测 好你",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringReverse(tt.args.src); got != tt.want {
				t.Errorf("StringReverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

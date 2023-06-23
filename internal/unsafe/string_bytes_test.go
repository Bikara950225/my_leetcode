package unsafe

import (
	"reflect"
	"testing"
)

func TestStringToByteUnsafe(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "happy path",
			args: args{
				src: "Hello World; 你好世界",
			},
			want: []byte("Hello World; 你好世界"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToByteUnsafe(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToByteUnsafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToByteUnsafe2(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "happy path",
			args: args{
				src: "Hello World; 你好世界",
			},
			want: []byte("Hello World; 你好世界"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StringToByteUnsafe2(tt.args.src)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToByteUnsafe2() = %v, want %v", got, tt.want)
			}
		})
	}
}

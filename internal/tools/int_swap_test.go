package tools

import "testing"

func TestIntSwap(t *testing.T) {
	mockI := int(^uint(0) >> 1)

	type args struct {
		src1F func() *int
		src2F func() *int
	}
	tests := []struct {
		name string
		args args
		wantSrc1 int
		wantSrc2 int
	}{
		{
			name: "swap",
			args: args{
				src1F: func() *int {
					i := 1
					return &i
				},
				src2F: func() *int {
					i := 2
					return &i
				},
			},
			wantSrc1: 2,
			wantSrc2: 1,
		},
		{
			name: "src1 is nil",
			args: args{
				src1F: func() *int {
					return nil
				},
				src2F: func() *int {
					i := 2
					return &i
				},
			},
			wantSrc2: 2,
		},
		{
			name: "src2 is nil",
			args: args{
				src1F: func() *int {
					i := 1
					return &i
				},
				src2F: func() *int {
					return nil
				},
			},
			wantSrc1: 1,
		},
		{
			name: "src1 == src2",
			args: args{
				src1F: func() *int {
					return &mockI
				},
				src2F: func() *int {
					return &mockI
				},
			},
			wantSrc1: mockI,
			wantSrc2: mockI,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			src1 := tt.args.src1F()
			src2 := tt.args.src2F()

			IntSwap(src1, src2)
			if src1 != nil && *src1 != tt.wantSrc1 {
				t.Errorf("TestIntSwap() error, src1 not expect result; result: %d, expect: %d",
					*src1, tt.wantSrc1,
				)
				return
			}
			if src2 != nil && *src2 != tt.wantSrc2 {
				t.Errorf("TestIntSwap() error, src2 not expect result; result: %d, expect: %d",
					*src2, tt.wantSrc2,
				)
				return
			}
		})
	}
}

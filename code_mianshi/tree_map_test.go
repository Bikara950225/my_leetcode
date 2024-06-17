package code_mianshi

import "testing"

func TestBtreeMapDemo(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "happy path",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			treeMapDemo()
		})
	}
}

func Test_bTreeMapDemo(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "happy path",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bTreeMapDemo()
		})
	}
}

func Test_bTreeMapSparseIndexDemo(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "happy path",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bTreeMapSparseIndexDemo()
		})
	}
}

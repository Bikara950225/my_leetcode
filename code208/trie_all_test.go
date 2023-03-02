package main

import (
	"reflect"
	"testing"
)

func Test_trieAll_Insert_Search(t *testing.T) {
	type fields struct {
		words []string
	}
	type args struct {
		words []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []bool
	}{
		{
			name: "happy path",
			fields: fields{
				words: []string{
					"/object/123",
					"/object/12/47",
					"你好",
				},
			},
			args: args{
				words: []string{
					"/object/123",
					"/object/12",
					"/object/12/47",
					"/123",
					"你好",
					"你好123",
				},
			},
			want: []bool{
				true, false, true, false,
				true, false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewTrieAll()
			for _, w := range tt.fields.words {
				s.Insert(w)
			}
			var ret []bool
			for _, w := range tt.args.words {
				ret = append(ret, s.Search(w))
			}
			if !reflect.DeepEqual(ret, tt.want) {
				t.Errorf("Test_trieAll_Insert_Search() error, not expect: %v", ret)
				return
			}
		})
	}
}

func Test_trieAll_StartsWith(t *testing.T) {
	type fields struct {
		words []string
	}
	type args struct {
		prefix []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []bool
	}{
		{
			name: "happy path",
			fields: fields{
				words: []string{
					"/object/123",
					"你好",
				},
			},
			args: args{
				prefix: []string{
					"/object/123",
					"/object/12",
					"/123",
					"你好",
					"你好123",
					"你",
				},
			},
			want: []bool{
				true, true, false, true, false, true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewTrieAll()
			for _, w := range tt.fields.words {
				s.Insert(w)
			}
			ret := []bool{}
			for _, w := range tt.args.prefix {
				ret = append(ret, s.StartsWith(w))
			}
			if !reflect.DeepEqual(ret, tt.want) {
				t.Errorf("Test_trieAll_StartsWith() error, not expect: %v", ret)
				return
			}
		})
	}
}

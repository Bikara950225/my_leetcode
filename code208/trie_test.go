package main

import (
	"reflect"
	"testing"
)

func Test_trie_Insert_Search(t *testing.T) {
	type fields struct {
		word []string
	}
	type args struct {
		word []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []bool
	}{
		{
			name: "test",
			fields: fields{
				word: []string{
					"ab", "abc", "bc",
				},
			},
			args: args{
				word: []string{
					"ab", "abd", "bc", "bca",
				},
			},
			want: []bool{
				true, false, true, false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Constructor()
			for _, w := range tt.fields.word {
				s.Insert(w)
			}

			gotlist := make([]bool, 0, len(tt.args.word))
			for _, w := range tt.args.word {
				gotlist = append(gotlist, s.Search(w))
			}
			if !reflect.DeepEqual(gotlist, tt.want) {
				t.Errorf("Test_trie_Insert_Search() error, got: %v", gotlist)
				return
			}
		})
	}
}

package main

import (
	"testing"

	"github.com/YoungGoofy/wb_l2/develop/dev02/unpack"
)

func TestUnpack(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		want  string
	}{
		{
			desc:  "normal",
			input: "a4df5cvs1",
			want:  "aaaadfffffcvs",
		},
		{
			desc:  "no numbers, only letters",
			input: "abcd",
			want:  "abcd",
		},
		{
			desc:  "only numbers",
			input: "45",
			want:  "",
		},
		{
			desc:  "empty string",
			input: "",
			want:  "",
		},
		{
			desc:  "multiple digit number",
			input: "a9df5cvs1",
			want:  "aaaaaaaaadfffffcvs",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			res, err := unpack.Unpack(tc.input)
			if err != nil {
				t.Error(err)
			}
			if res != tc.want {
				t.Errorf("Got: %s\nWant: %s", res, tc.want)
			}
		})
	}
}

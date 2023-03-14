package main

import "testing"

func TestAnagram(t *testing.T) {
	testCases := []struct {
		desc string
		data []string
		want map[string][]string
	}{
		{
			desc: "normal",
			data: []string{
				"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "столик", "листок",
			},
			want: map[string][]string{
				sortString("листок"): {"листок", "слиток", "столик"},
				sortString("пятак"):  {"пятак", "пятка", "тяпка"},
			},
		},
		{
			desc: "none",
			data: []string{
				"123f", "ghskjas", "sfasdfk",
			},
			want: map[string][]string{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			res := Anagram(&tc.data)
			if res != &tc.want {
				t.Errorf("Want %v, got %v", tc.want, *res)
			}
		})
	}
}

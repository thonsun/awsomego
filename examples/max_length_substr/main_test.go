package main

import (
	"testing"
)

func TestFindMaxLengthSubString(t *testing.T) {
	var data = []struct{
		in string
		expected int
	}{
		{"abcdef",6},
		{"abccd",3},
		{"aaaaa",1},
		{"abababa",2},
	}

	for _,d := range data{
		if actual := FindMaxLengthSubString([]rune(d.in));actual != d.expected {
			t.Error("[+]error: ",d.in,d.expected,actual)
		}
	}
}

package _51_reverse_words

import "testing"

func TestReverseWords(t *testing.T) {
	var data  = []struct {
		in string
		expect string
	}{
		{"the sky is blue","blue is sky the"},
		{"    hello world","world hello"},
	}

	for _,d := range data{
		if out := ReverseWords(d.in); out != d.expect{
			t.Error("+error: ",d.in," | ",d.expect," | ",out)
		}
	}
}

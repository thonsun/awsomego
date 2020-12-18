package _67_two_nums_mul

import (
	"testing"
)

func TestReverse(t *testing.T) {
	var data = []struct{
		in string
		expect string
	}{
		{"123456","654321"},
		{"123","321"},
		{"1","1"},
	}

	for _,d := range data{
		if out := reverse(d.in);out!=d.expect{
			t.Error("!" + d.in + "|" + d.expect +"|" + out)
		}
	}
}

func TestBigNumbersMul(t *testing.T) {
	var data = []struct{
		a string
		b string
		expect string
	}{
		{"15","15","225"},
		{"0","100","0"},
	}

	for _,d := range data{
		if out := BigNumbersMul(d.a,d.b);out != d.expect {
			t.Error("!",d.a,d.b,d.expect,out)
		}
	}
}

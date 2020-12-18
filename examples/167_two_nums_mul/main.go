package _67_two_nums_mul

import "strconv"

func BigNumbersMul(num1 string, num2 string) (result string){
	if num1[0] - '0' == 0 || num2[0] - '0' == 0 {
		return "0"
	}

	//a := reverse(num1)
	//b := reverse(num2)

	a := num1
	b := num2

	var ret = make([]int,len(a)+len(b))
	for i:=0;i<len(a);i++{
		x := a[i]-'0'
		for j:=0;j<len(b);j++{
			y := b[j]-'0'
			ret[i+j] += (ret[i+j+1]+int(x*y))/10
			ret[i+j+1] = (ret[i+j+1]+int(x*y))%10
		}
	}

	result = ""
	for i:=0;i<len(ret);i++{
		if i ==0 && ret[i] == 0{
			continue
		}
		result += strconv.Itoa(ret[i])
	}
	return
}

func reverse(s string) string {
	r := []rune(s)
	for i,j := 0,len(r)-1;i<j;i,j = i + 1,j -1 {
		r[i],r[j] = r[j],r[i]
	}
	return string(r)
}

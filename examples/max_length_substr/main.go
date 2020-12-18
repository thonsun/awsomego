package main

import "fmt"

func main() {
	
}

func FindMaxLengthSubString(str []rune) int {
	var max_substr string = ""
	var tmp_substr = make([]rune,0)
	var max_length int = 0
	var length int = 0
	var refer = make(map[rune]bool)

	for i:=0;i<len(str);i++{
		if _,ok := refer[str[i]];ok{
			if length >= max_length {
				max_length = length
				max_substr = string(tmp_substr)
			}
			length = 0
			refer = make(map[rune]bool)
			tmp_substr = make([]rune,0)
		}

		length += 1
		refer[str[i]] = true
		tmp_substr = append(tmp_substr,str[i])
	}

	if length >= max_length {
		max_length = length
		max_substr = string(tmp_substr)
	}
	fmt.Println(max_length,":",string(max_substr))
	return max_length
}

package _51_reverse_words

func ReverseWords(smt string) string{
	array := []string{}
	words := []rune(smt)
	var start int

	step := -1
	for i:=0;i<len(words);i++{
		if step == -1 && words[i] == []rune(" ")[0] {
			continue
		}
		if step == -1 && words[i] != []rune(" ")[0] {
			step = 1
			start = i
		}
		if step == 1 && words[i] == []rune(" ")[0] {
			array = append(array,string(words[start:i]))
			step = -1
		}
	}
	if words[len(words)-1] != []rune(" ")[0] {
		array = append(array,string(words[start:]))
	}

	res := ""
	for i:= len(array)-1; i>= 0; i-- {
		if i == len(array) -1 {
			res = array[i]
		}else{
			res = res + " " + array[i]
		}

	}
	return res
}

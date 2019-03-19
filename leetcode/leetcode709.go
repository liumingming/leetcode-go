package leetcode

func toLowerCase(str string) string {
	length := len(str)
	var res []rune
	for i:=0; i<length; i++ {
		c := str[i]
		if c >= 'A' && c <= 'Z' {
			res = append(res ,rune(c+32))
		}else {
			res = append(res ,rune(c))
		}
	}
	return string(res)
}

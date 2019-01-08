package leetcode

//反转字符串
func reverseString(s string) string {
	b  := []byte(s)
	for i := 0; i< len(s)/2; i++ {
		last := len(s)-1-i
		b[i] = s[last]
		b[last] = s[i]
	}
	return string(b)
}
package leetcode

func reverseVowels(s string) string {
	left := 0
	right := len(s)-1
	var res = []byte(s)
	for left < right {
		for !isVoe(s[left]) && left < right {
			left+=1
		}
		for !isVoe(s[right]) && left < right {
			right-=1
		}

		if isVoe(s[left]) && isVoe(s[right]) {
			res[left], res[right] = res[right], res[left]
			left+=1
			right-=1
		}
	}
	return string(res)
}

func isVoe(ch byte) bool {
	voes := []byte{'a', 'e', 'i', 'o', 'u'}
	for _, v := range voes {
		if ch == v {
			return true
		}
	}
	return false
}

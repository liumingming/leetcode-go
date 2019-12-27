package leetcode


//problem-14 最长公共前缀
func longestCommonPrefix(strs []string) string {
	var curr string
	if len(strs) == 0 || len(strs[0]) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	for i := 1; i < len(strs); i++ {
		if len(strs[i]) == 0 {
			return curr
		}
		var temp string
		for j := 1; j <= len(strs[i]); j++ {
			if len(strs[0]) < j || strs[i][0:j] != strs[0][0:j] {
				break
			}
			temp =  strs[0][0:j]
		}

		if len(temp) == 0 {
			break
		}

		if curr == "" || len(curr) > len(temp) {
			curr = temp
		}
	}
	return curr
}



















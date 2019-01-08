package leetcode


func lengthOfLongestSubstring(s string) int {
	max := 1
	for i := 0; i < len(s) ; i++  {
		count := 1
		for j := i+1; j< len(s) ; j++  {
			if s[i] == s[j] {
				break
			}
			count++
		}

		if count > max {
			max = count
		}

	}

	return max
}

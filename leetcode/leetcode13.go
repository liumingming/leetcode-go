package leetcode

//leetcode problem-13 罗马数字转整数
func romanToInt(s string) int {
	roman_map := make(map[string]int)
	roman_map = map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
	}
	sum := 0

	for i:= 0; i < len(s); i++ {
		if  i< len(s)-1 && roman_map[string(s[i])] < roman_map[string(s[i+1])] {
			sum += roman_map[string(s[i+1])] - roman_map[string(s[i])]
			i++
		} else {
			sum += roman_map[string(string(s[i]))]
		}
	}

	return sum
}


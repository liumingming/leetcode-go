package leetcode

func numJewelsInStones(J string, S string) int {
	Jmap := make(map[rune]int)
	for i := 0; i<len(J); i++ {
		Jmap[rune(J[i])] = 1
	}
	var count int
	for i := 0; i < len(S); i++ {
		_, ok := Jmap[rune(S[i])]
		if ok {
			count ++
		}
	}
	return count
}
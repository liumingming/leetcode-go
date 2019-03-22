package leetcode

func uniqueMorseRepresentations(words []string) int {
	var count int
	var dict  = []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
	var moseMap = make(map[string]int)
	for _,  word := range words {
		var moseStr string
		for _, ch := range  word {
			pos := ch - 97
			if pos < 0 || pos > 25 {
				break
			}
			moseStr += dict[pos]
		}
		if _, ok := moseMap[moseStr]; !ok {
			count++
			moseMap[moseStr] = 1
		}

	}
	return count
}


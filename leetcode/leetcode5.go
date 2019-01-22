package leetcode

const SIZE = 1000

func Manacher(str []byte) string {

	manacher := ""

	if len(str) == 0 {
		return manacher
	}

	var result [SIZE][SIZE]bool

	for i, j := 0, 0; j < len(str); j++ {
		i = j

		for i >= 0 {
			if str[i] == str[j] && ((j-i) < 2 || result[i+1][j-1]) {
				result[i][j] = true
				if j-i+1 > len(manacher) {
					manacher = string(str[i : j+1])
				}
			}

			i--
		}
	}

	return manacher
}

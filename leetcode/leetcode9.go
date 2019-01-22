package leetcode

import (
	"math"
	"strconv"
)

//leetcode problem-9 回文数
func isPalindrome(x int) bool {
	tmp := x
	if x < 0 {
		return false
	}

	var sum int
	x = int(math.Abs(float64(x)))
	var count = len(strconv.Itoa(x))

	for {
		sum += (x % 10)  * int(math.Pow10(count-1))
		if x / 10 == 0 {
			break
		}
		x = x / 10
		count--
	}

	if tmp == sum {

		return  true
	}

	return false
}


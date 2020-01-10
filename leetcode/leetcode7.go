package leetcode

import (
	"math"
	"strconv"
)

//problem-7 整数反转
func reverse(x int) int {
	var sum int
	var flag int
	if x < 0 {
		flag = -1
	}else {
		flag = 1
	}

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

	if sum >= math.MaxInt32-1 || sum <= math.MinInt32-1 {
		return 0
	}

	return int(sum) * flag
}
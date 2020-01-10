package seven

import "math/rand"


/*
7.4有一个函数func1能返回0和1两个值，返回0和1的概率都是1/2，问怎么利用这个函数得到另一个函数func2
使得func也只能返回0和1，且返回的概率为1/4
 */
func func1() int {
	return rand.Intn(2)
}

func func2() int  {
	a1 := func1()
	a2 := func1()

	tmp := a1
	tmp |= a2 << 1

	if tmp == 0 {
		return 0
	} else {
		return 1
	}
}



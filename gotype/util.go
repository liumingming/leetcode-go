package gotype

func CreateNode(node *ListNode, max int)  {
	cur := node
	for i := 1; i < max ; i++  {
		cur.Next = &ListNode{}
		cur.Next.Data = i
		cur = cur.Next
	}
}

func Abs(x int) int  {
	switch  {
	case x < 0:
		return -x
	case x == 0:
		return 0
	}
	return x
}

func Min(a, b int) int  {
	if a < b {
		return a
	} else {
		return b
	}
}


func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min3(a, b, c int) int  {
	min := Min(a, b)
	return Min(min, c)
}

func Max3(a, b, c int) int  {
	max := Max(a, b)
	return Max(max, c)
}

func IsNumber(c rune) bool  {
	return c >= '0' && c <= '9'
}

func SwapInt(data []int, x, y int)  {
	tmp := data[x]
	data[x] = data[y]
	data[y] = tmp
}

func SwapRune(data []rune, x, y int)  {
	tmp := data[x]
	data[x] = data[y]
	data[y] = tmp
}

func IsOne(n, i int) bool {
	return (uint(n) & (uint(1) << uint(i))) == 1
}
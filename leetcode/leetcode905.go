package leetcode

//problem 905
func sortArrayByParity(A []int)  []int  {
	ret := make([]int, len(A))

	i := 0
	j := 1

	for _, n := range A {
		if n % 2 == 0 {
			ret[i] = n
			i++
		} else {
			ret[len(A)-j] = n
			j++
		}
	}
	return  ret
}
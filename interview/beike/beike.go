package main

import "fmt"

//
//var res [][]int
//
//var arr = [][]int{
//	{1,2,3},
//	{4,5,6},
//}
//
//func generateArr() [][]int {
//	res = make([][]int, 0)
//
//
//	for i:= 0; i < len(arr); i++ {
//		helper(arr[i], arr[i+1])
//	}
//
//
//	return res
//}
//
//func helper(arr1 []int, arr2 []int ) [][]int {
//	var r [][]int
//	for i:=0; i < len(arr1) ; i++ {
//		for j:=0; j<len(arr2); j++ {
//			var tmp = make([]int,0)
//			tmp = append(tmp, arr1[i], arr2[j])
//			r = append(r, tmp)
//		}
//	}
//	return r
//}
//
//func main()  {
//	result helper(arr[0], arr[1])
//}

func main() {
	var input = [][]int{
		{1, 2, 3,4},
		{5, 6, 7,8},
		{9, 10,11,12},
		{13,14,15,16},
	}
	output := helper(input, nil, 0)

	fmt.Println(output)

}


func helper(arr [][]int, temp []int, begin int) [][]int {
	var output [][]int
	if len(arr) <= begin+1 {
		for _, val := range arr[begin] {
			var t = append(temp, val)
			output = append(output, t)
		}
		return output
	}

	for _, val := range arr[begin] {
		var t = append(temp, val)
		for _, subArr := range helper(arr, t, begin+1) {
			output = append(output, subArr)
		}
	}
	return output
}

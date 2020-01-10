package main

import (
	"fmt"
)

/*组合,不把四百道题刷得孰能生巧，我坚决不去面试*/
var arr [][]int = [][]int{
	{1,2,4},
	{2,3,4},
	{3,4,5},
}

func PrintArr(arr [][]int, row int)  {
	var s [][]int
	
	for {
		for i:=0; i<row ;i++  {
			var k []int

			k = append(k, arr[i][step])
			s = append(s, k)
		}
		
		break
	}
	

	fmt.Println(s)
}

func generArr(arr1, arr2 []int)  {

}

func main()  {
	PrintArr(arr, len(arr))
}


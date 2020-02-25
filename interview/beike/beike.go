package main

import "fmt"

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

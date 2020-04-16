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


var m = map[int]string{
	1: "一",
	2: "二",
	3: "三",
	4: "四",
	5: "五",
	6: "六",
	7: "七",
	8: "八",
	9: "九",
	0: "零",
}

func translate(num int) string {
	if num > 100000 || num < 0 {
		return ""
	}

	var res []string
	var remind int
	var div int
	remind = num / 10000
	div  = num % 10000
	if remind > 0 {
		r, ok := m[remind]; if ok {
			res = append(res, r, "万")
		}
	}

	remind = div / 1000
	div  = num % 1000
	if remind > 0 {
		r, ok := m[remind]; if ok {
			res = append(res, r, "千")
		}
	} else {
		res = append(res, "零")
	}

	remind = div / 100
	div  = num % 100
	if remind > 0 {
		r, ok := m[remind]; if ok {
			res = append(res, r, "百")
		}
	}else {
		res = append(res, "零")
	}

	remind = div / 10
	div  = num % 100
	if remind > 0 {
		r, ok := m[remind]; if ok {
			res = append(res, r, "十")
		}
	} else {
		res = append(res, "零")
	}

	remind = num%10
	if remind > 0 {
		r, ok := m[remind]; if ok {
			res = append(res, r,)
		}
	}

	var temp string
	var left int
	var right int = len(res) -1
	for res[left] == "零" {
		left++
	}

	for res[right] == "零" {
		right--
	}
	for left <= right {
		str := res[left]

		if str == "零" && res[left-1]== "零"{
			left++
			continue
		}
		left++
		temp += str
	}

	return temp
}
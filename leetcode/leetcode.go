package leetcode

import (
	"math"
	"strconv"
)


//leetcode problem-1 两数之和
func twoSum(nums []int, target int) []int {
	var (
		left int
		nums_map = make(map[int]int)
	)
	for index, num:= range nums {
		left = target - num
		value, ok := nums_map[left]
		if ok {
			return []int{index, value}
		} else {
			nums_map[num] = index
		}
	}
	return  []int{}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//type ListNode struct {
//	    Val int
//	 Next *ListNode
//}
//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	root := &ListNode{}
//	for {
//
//		if l1.Next == nil {
//			break
//		}
//
//		if l2.Next == nil {
//			break
//		}
//
//		value := l1.Val + l2.Val
//
//		if value >= 10 {
//
//		}
//
//	}
//}


//leetcode problem-5 最长回文子串
var pmap =  make(map[string]bool)
func longestPalindrome(s string) string {
	length := len(s)

	if length <= 1 {
		return s
	}

	isp, ok := pmap[s]
	if ok{
		if isp {
			return s
		} else {
			return ""
		}
	}
	reverse := reverseString(s)

	if s == reverse {
		pmap[s] = true
		return s
	} else {
		pmap[s] = false
	}

	res1 := longestPalindrome(s[0:length-1])

	res2 :=  longestPalindrome(s[1:length])

	if len(res1) >= len(res2)  {
		return res1
	} else {
		return res2
	}
}


//leetcode problem-7 整数反转
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


//leetcode problem-13 罗马数字转整数
func romanToInt(s string) int {
	roman_map := make(map[string]int)
	roman_map = map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
	}
	sum := 0

	for i:= 0; i < len(s); i++ {
		if  i< len(s)-1 && roman_map[string(s[i])] < roman_map[string(s[i+1])] {
			sum += roman_map[string(s[i+1])] - roman_map[string(s[i])]
			i++
		} else {
			sum += roman_map[string(string(s[i]))]
		}
	}

	return sum
}


//leetcode problem-14 最长公共前缀
func longestCommonPrefix(strs []string) string {
	var curr string
	if len(strs) == 0 || len(strs[0]) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	for i := 1; i < len(strs); i++ {
		if len(strs[i]) == 0 {
			return curr
		}
		var temp string
		for j := 1; j <= len(strs[i]); j++ {
			if len(strs[0]) < j || strs[i][0:j] != strs[0][0:j] {
				break
			}
			temp =  strs[0][0:j]
		}

		if len(temp) == 0 {
			break
		}

		if curr == "" || len(curr) > len(temp) {
			curr = temp
		}
	}
	return curr
}


//leetcode problem-26 删除有序数组中的重复元素
func removeDuplicates(nums []int) int {
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}

	}
	return i+1
}











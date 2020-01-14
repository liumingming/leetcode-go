package leetcode

//problem-1 两数之和
func twoSum(nums []int, target int) []int {
	var numsMaps = make(map[int]int)
	for index, num:= range nums {
		left := target - num
		value, ok := numsMaps[left]
		if ok {
			return []int{index, value}
		}

		numsMaps[num] = index
	}
	return  []int{}
}


package leetcode

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


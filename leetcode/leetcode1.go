package leetcode

//leetcode problem-1 两数之和
func twoSum(nums []int, target int) []int {
	var (
		left     int
		numsMaps = make(map[int]int)
	)
	for index, num:= range nums {
		left = target - num
		value, ok := numsMaps[left]
		if ok {
			return []int{index, value}
		} else {
			numsMaps[num] = index
		}
	}
	return  []int{}
}


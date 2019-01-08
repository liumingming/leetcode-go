package leetcode

import "sort"

//leetcode problem-136 只出现一次的数字1
func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	sort.Ints(nums)

	for i := 0; i<len(nums)-1; i++ {
		if i == 0 && nums[i] != nums[i+1] {
			return nums[i]
		}

		if nums[i] != nums[i+1] && nums[i] != nums[i-1] {
			return nums[i]
		}
	}

	return nums[len(nums)-1]
}

package leetcode

import "fmt"

//leetcode problem-80 删除有序数组中的重复元素, 每个元素最多出现两次
func RemoveDuplicates2(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	index := 2
	for j := 2; j < len(nums); j++ {
		if nums[j] == nums[index-2]  {
			index++
			nums[j] = nums[index]
		}
	}
	fmt.Println(nums)
	return index+1
}
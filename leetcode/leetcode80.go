package leetcode

//problem-80 删除有序数组中的重复元素, 每个元素最多出现两次
func RemoveDuplicates2(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	index := 0
	for i := 2; i < len(nums); i++ {
		if nums[index] == nums[i] {
			nums[index] = nums[i+1]
		}
		index++
	}
	return index+1
}
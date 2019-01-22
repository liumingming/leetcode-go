package leetcode

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
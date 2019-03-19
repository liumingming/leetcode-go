package basic

/*冒泡排序*/
func bubbleSort(nums []int) []int {
	length := len(nums)
	for i:= 0; i< length; i++{
		for j:= i+1; j < length ;j++  {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

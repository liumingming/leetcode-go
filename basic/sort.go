package basic

/*冒泡排序*/
func BubbleSort(nums []int) []int {
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


/*插入排序*/
func InsertSort(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}
	for i := 1; i < length; i++ {
		for j := 0; j < i; j++ {
			if nums[j] > nums[i] {
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
	}
	return nums
}

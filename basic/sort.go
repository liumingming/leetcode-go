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

/*快速排序*/
func QuickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	var pivot = nums[0]
	var less []int
	var greater []int
	for i := 1 ; i < len(nums) ;i++  {
		if nums[i] <= pivot {
			less = append(less, nums[i])
		}

		if nums[i] > pivot {
			greater = append(greater, nums[i])
		}

	}
	return append(append(QuickSort(less),  pivot), QuickSort(greater)...)
}

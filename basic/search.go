package basic

//二分查找
func binarySearch(nums []int, target int) int {
	low := 0
	high := len(nums)
	mid := (low + high) / 2

	for low < high {
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid + 1
		} else {
			low = mid + 1
		}
		mid = (low + high) / 2
	}

	return -1
}
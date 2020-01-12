package leetcode

var result = [][]int{}

func subsets(nums []int) [][]int {
	var item  []int

	result = append(result, item)

	generate(0, nums, item)

	return result

}

func generate(i int,  nums []int, item []int){
	if i >= len(nums) {
		return
	}

	item = append(item, nums[i])

	result = append(result, item)

	generate(i+1, nums, item)

	var temp = make([]int, len(item)-1)

	copy(temp,item[:len(item)-1])

	generate(i+1, nums, temp)
}

package basic

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	nums := []int{4,1,5,3,2,6}
	res := BubbleSort(nums)
	//targets := []int{1,2,3,4,5,6}
	fmt.Println(res)
}


func TestInsertSort(t *testing.T) {
	nums := []int{4,1,5,3,2,6}
	res := InsertSort(nums)
	//targets := []int{1,2,3,4,5,6}
	fmt.Println(res)
}


func TestQuickSort(t *testing.T) {
	nums := []int{4,1,5,3,2,6}
	res := QuickSort(nums)
	//targets := []int{1,2,3,4,5,6}
	fmt.Println(res)
}
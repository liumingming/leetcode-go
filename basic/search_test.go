package basic

import (
	"testing"
)

func TestBinerySearch(t *testing.T) {
	nums := []int{0,1,2,3,4,5,6}
	res := binarySearch(nums, 5)
	if res != 5 {
		t.Errorf("res is %d and target is %d", res, 5)
	}

	res = binarySearch(nums, 10)
	if res != -1 {
		t.Errorf("res is %d and target is %d", res, 10)
	}
}


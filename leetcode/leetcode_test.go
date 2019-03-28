package leetcode

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	input := "minglming"
	res := reverseString(input)
	if res != "gnimlgnim" {
		t.Errorf(`reverseString(%q) = %q`, input,  res)
	}
}


func TestTwoSum(t *testing.T) {
	input := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(input, target)
	if res[0] == 0 && res[1] == 1 {
		t.Errorf(`twoSum(%q) = %q`, input,  res)
	}
}



func TestSigerNumber(t *testing.T) {
	input := []int{4,1,2,1,2}
	res := singleNumber(input)
	if res != 4 {
		t.Errorf(`twoSum(%q) = %d`, input,  res)
	}
}

func TestRomanToInt(t *testing.T) {
	input :=  "MCMXCIV"
	res := romanToInt(input)

	if res != 1994 {
		t.Errorf(`romanToInt(%q) = %d`, input,  res)
	}

}



func TestReverse(t *testing.T) {
	a := 321

	b := reverse(a)

	if b != 123 {
		t.Errorf("b %d is not the reversed a %d", b, a)
	}



	c := -321

	d := reverse(c)

	if d != -123 {
		t.Errorf("d %d is not the reversed c %d", b, c)
	}


	e := 1534236469

	f := reverse(e)

	if f != 0 {
		t.Errorf("d %d is not the reversed c %d", e, f)
	}
}


func TestIsPalindrome(t *testing.T) {
	a := 121
	res := isPalindrome(a)
	if !res {
		t.Errorf("a %d is palindrome", a)
	}
}


func TestLongestCommonPrefix(t *testing.T) {
	//strs := []string{"flower","flow","flight"}
	//res := longestCommonPrefix(strs)
	//fmt.Println(res)
	//if res == "" {
	//	t.Errorf("res %s is fl", res)
	//}


	//strs := []string{"c","aaa","ccc"}
	//res := longestCommonPrefix(strs)
	//if res != "" {
	//	t.Errorf("res %s is", res)
	//}


	//strs := []string{"c","c"}
	//res := longestCommonPrefix(strs)
	//fmt.Println(res)
	//if res != "" {
	//	t.Errorf("res %s is", res)
	//}

	strs := []string{"c","c", "b"}
	res := longestCommonPrefix(strs)
	fmt.Println(res)
	if res != "" {
		t.Errorf("res %s is", res)
	}
}


func TestLongestP(t *testing.T) {
	//str := "babaddtattarrattatddetartrateedredividerb"
	//res := longestPalindrome(str)
	//
	//if res != "ddtattarrattatdd" {
	//	t.Errorf("res is %s and failed", res)
	//}
	str := "eabcb"
	res := longestPalindrome(str)

	if res != "bcb" {
		t.Errorf("res is %s and failed", res)
	} else {
		fmt.Println(res)
	}
}

func TestString(t *testing.T) {
	str := "babad"
	fmt.Print(str[0:len(str)-1])
}

//反转字符串
func reverseString(s string) string {
	b  := []byte(s)
	for i := 0; i< len(s)/2; i++ {
		last := len(s)-1-i
		b[i] = s[last]
		b[last] = s[i]
	}
	return string(b)
}

func Test_TolowerCase(t *testing.T)  {
	 str :=  "ALOVELY"
	 target  := "alovely"

	 res := toLowerCase(str)
	 if res != target {
		 t.Errorf("res is %s and target is %s", res, target)
	 }

}

func Test_NumJewelsInStones(t *testing.T)  {
	J :=  "aA"
	S :=  "aAAbbbb"
	res := numJewelsInStones(J, S)
	if res != 3 {
		t.Errorf("res is %d and target is %d", res, 3)
	}

}

func Test_SortedSquares(t *testing.T)  {
	A :=  []int{-4,-1,0,3,10}
	res := sortedSquares(A)
	fmt.Println(res)
}



func Test_UniqueMorseRepresentations(t *testing.T) {
	words := []string{"gin", "zen", "gig", "msg"}
	res := uniqueMorseRepresentations(words)
	fmt.Println(res)
}

func Test_GenerateMatrix(t *testing.T) {
	res := generateMatrix(3)
	fmt.Println(res)
}
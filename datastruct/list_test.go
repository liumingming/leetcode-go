package datastruct

import (
	"fmt"
	"testing"
)

func TestNewSingleList(t *testing.T) {
	 single := NewSingleList()
	 single.AddNode(1)
	 single.AddNode(2)
	 single.AddNode(3)
	 single.AddNode(4)
	 single.AddNode(5)
	 single.Print()
}



func TestSingleList_Reverse(t *testing.T) {
	single := NewSingleList()
	single.AddNode(1)
	single.AddNode(2)
	single.AddNode(3)
	single.AddNode(4)
	single.AddNode(5)
	single.Print()
	single.Reverse()
	single.Print()
}

func TestArr(t *testing.T)  {
	var m  = make([]int, 10, 10)

	for i:=0; i<10 ;i++  {
		fmt.Println(m[i])
	}
}

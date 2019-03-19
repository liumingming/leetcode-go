package basic

import "testing"

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

package gotype

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	head := &ListNode{}
	CreateList(head, 8)
	PrintList(head)
	fmt.Println()
	Reverse(head)
	PrintList(head)
}


func TestReverseNoHeadNode(t *testing.T) {

	head := CreateListNoHead()
	PrintListWithNoHead(head)
	fmt.Println()
	ReverseNoHeadNode(head)
	PrintListWithNoHead(head)
}



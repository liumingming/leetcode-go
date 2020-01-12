package gotype

import (
	"fmt"
	"math/rand"
)

type ListNode struct {
	Data interface{}
	Next *ListNode
}

func NewListNode()  *ListNode {
	return &ListNode{}
}

func NewListNodeWithData(data int)  *ListNode {
	return &ListNode{Data:data}
}

func CreateList(head *ListNode, num int)  {
	for i:=0;i<num;i++  {
		node := NewListNode()
		data := rand.Intn(num)
		node.Data = data
		node.Next = head.Next
		head.Next = node
	}
}

func CreateListNoHead() *ListNode {
	node1 := NewListNodeWithData(1)
	node2 := NewListNodeWithData(2)
	node1.Next = node2
	node3 := NewListNodeWithData(3)
	node2.Next = node3
	node4 := NewListNodeWithData(4)
	node3.Next = node4
	node5 := NewListNodeWithData(5)
	node4.Next = node5
	return node1
}

func PrintList(head *ListNode)  {
	for head.Next != nil  {
		fmt.Print(head.Next.Data, " ")
		head = head.Next
	}
}

func PrintListWithNoHead(head *ListNode)  {
	for head != nil {
		fmt.Print(head.Data, " ")
		head = head.Next
	}
}

type BTreeNode struct {
	Data interface{}
	LeftChild *BTreeNode
	RightChild *BTreeNode
}

func NewBTreeNode() *BTreeNode  {
	return &BTreeNode{}
}

type TrieNode struct {
	IsLeaf bool
	URL string
	Child []*TrieNode
}

func NewTrieNode(count int) *TrieNode  {
	return &TrieNode{IsLeaf:false, URL:"",Child:make([]*TrieNode, count)}
}

type AVLNode struct {
	Data int
	Height int
	Count int
	Left *AVLNode
	Right *AVLNode
}

func NewAVLNode(data int) *AVLNode  {
	return &AVLNode{
		Data:   data,
		Height: 1,
		Count:  1,
		Left:   nil,
		Right:  nil,
	}
}


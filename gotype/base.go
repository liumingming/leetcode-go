package gotype

type ListNode struct {
	Data interface{}
	Next *LNode
}

func NewListNode()  *ListNode {
	return &ListNode{}
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

